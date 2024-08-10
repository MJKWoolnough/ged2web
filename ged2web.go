// Ged2Web converts a GEDCOM file into a SPA webpage that can be simply added to any existing website
package main // import "vimagination.zapto.org/ged2web"

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"vimagination.zapto.org/gedcom"
	"vimagination.zapto.org/rwcount"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type idMap map[gedcom.Xref]uint64

func (i idMap) GetID(ref gedcom.Xref) uint64 {
	id, ok := i[ref]
	if !ok {
		id = uint64(len(i)) + 1
		i[ref] = id
	}

	return id
}

type data []string

type gedcomData []data

func (g *gedcomData) Set(id uint64, data data) {
	if min := id + 1; min >= uint64(len(*g)) {
		if min >= uint64(cap(*g)) {
			h := make(gedcomData, min, min*2)

			copy(h, *g)

			*g = h
		} else {
			*g = (*g)[:min]
		}
	}

	(*g)[id] = data
}

func (g gedcomData) WriteTo(w io.Writer) {
	for n, d := range g {
		if n == 0 {
			io.WriteString(w, "[")
		} else {
			io.WriteString(w, ",[")
		}

		for m, p := range d {
			if m > 0 {
				io.WriteString(w, ",")
			}

			io.WriteString(w, p)
		}

		io.WriteString(w, "]")
	}
}

func run() error {
	var (
		err    error
		input  = flag.String("i", "-", "gedcom file")
		output = flag.String("o", "-", "output js file")
		html   = flag.Bool("h", false, "create full HTML file")
		module = flag.Bool("m", false, "just create gecom module")
		f      = os.Stdin
		w      = os.Stdout
	)

	flag.Parse()

	if *input != "-" {
		if f, err = os.Open(*input); err != nil {
			return fmt.Errorf("error opening input file (%s): %w", *input, err)
		}
	}

	indis, fams, err := processGedcom(f)

	f.Close()

	if err != nil {
		return err
	}

	if *output != "-" {
		if w, err = os.Create(*output); err != nil {
			return fmt.Errorf("error creating output file (%s): %w", *output, err)
		}
	}

	wr := &rwcount.Writer{Writer: w}

	if *html {
		io.WriteString(wr, htmlStart)
		io.WriteString(wr, jsStart)
	} else if *module {
		io.WriteString(wr, modStart)
	} else {
		io.WriteString(wr, jsStart)
	}

	indis.WriteTo(wr)

	if *module {
		io.WriteString(wr, modMid)
	} else {
		io.WriteString(wr, jsMid)
	}

	fams.WriteTo(wr)

	if *html {
		io.WriteString(wr, jsEnd)
		io.WriteString(wr, htmlEnd)
	} else if *module {
		io.WriteString(wr, modEnd)
	} else {
		io.WriteString(wr, jsEnd)
	}

	if wr.Err != nil {
		return fmt.Errorf("error writing to output file (%s): %w", *output, wr.Err)
	}

	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing output file (%s): %w", *output, err)
	}

	return nil
}

func processGedcom(f io.Reader) (gedcomData, gedcomData, error) {
	indiIDs := make(idMap)
	famIDs := make(idMap)
	indis := gedcomData{data{"", "", "", "", "0", "0"}}
	fams := gedcomData{data{"0", "0"}}
	r := gedcom.NewReader(f, gedcom.AllowMissingRequired, gedcom.IgnoreInvalidValue, gedcom.AllowUnknownCharset, gedcom.AllowTerminatorsInValue, gedcom.AllowWrongLength, gedcom.AllowInvalidEscape, gedcom.AllowInvalidChars)

	for {
		record, err := r.Record()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, nil, fmt.Errorf("error reading GEDCOM record: %w", err)
		}

		switch t := record.(type) {
		case *gedcom.Individual:
			person := append(make(data, 4, 6+len(t.SpouseOf)), "0", "0")

			if len(t.PersonalNameStructure) > 0 {
				name := strings.Split(string(t.PersonalNameStructure[0].NamePersonal), "/")

				var firstName, lastName string

				if t.Death.Date == "" {
					firstName = strings.Split(name[0], " ")[0]
				} else {
					firstName = strings.TrimSpace(name[0])
				}

				if len(name) > 1 {
					lastName = strings.TrimSpace(name[1])
				}

				person[0] = strconv.Quote(firstName)
				person[1] = strconv.Quote(lastName)
			}

			if t.Death.Date != "" {
				person[2] = strconv.Quote(strings.TrimSpace(string(t.Birth.Date)))
				person[3] = strconv.Quote(strings.TrimSpace(string(t.Death.Date)))
			}

			switch t.Gender {
			case "F", "f", "Female", "FEMALE", "female":
				person[4] = "2"
			case "M", "m", "Male", "MALE", "male":
				person[4] = "1"
			}

			if len(t.ChildOf) > 0 {
				person[5] = strconv.FormatUint(famIDs.GetID(t.ChildOf[0].ID), 10)
			}

			for _, spouse := range t.SpouseOf {
				person = append(person, strconv.FormatUint(famIDs.GetID(spouse.ID), 10))
			}

			indis.Set(indiIDs.GetID(t.ID), person)
		case *gedcom.Family:
			family := append(make(data, 0, 2+len(t.Children)), "0", "0")

			if t.Husband != "" {
				family[0] = strconv.FormatUint(indiIDs.GetID(t.Husband), 10)
			}

			if t.Wife != "" {
				family[1] = strconv.FormatUint(indiIDs.GetID(t.Wife), 10)
			}

			for _, child := range t.Children {
				family = append(family, strconv.FormatUint(indiIDs.GetID(child), 10))
			}

			fams.Set(famIDs.GetID(t.ID), family)
		}
	}

	return indis, fams, nil
}
