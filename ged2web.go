package main

import (
	"embed"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"vimagination.zapto.org/gedcom"
	"vimagination.zapto.org/javascript"
	"vimagination.zapto.org/jslib"
	"vimagination.zapto.org/parser"
)

var (
	//go:embed *.js lib/*.js
	files embed.FS

	ia, fa javascript.ArrayLiteral
	module = javascript.Module{
		ModuleListItems: []javascript.ModuleItem{
			{
				ExportDeclaration: &javascript.ExportDeclaration{
					Declaration: &javascript.Declaration{
						LexicalDeclaration: &javascript.LexicalDeclaration{
							LetOrConst: javascript.Const,
							BindingList: []javascript.LexicalBinding{
								{
									BindingIdentifier: token("people"),
									Initializer: &javascript.AssignmentExpression{
										ConditionalExpression: javascript.WrapConditional(&ia),
									},
								},
								{
									BindingIdentifier: token("families"),
									Initializer: &javascript.AssignmentExpression{
										ConditionalExpression: javascript.WrapConditional(&fa),
									},
								},
							},
						},
					},
				},
			},
		},
	}
)

func opener(url string) (*javascript.Module, error) {
	url = url[1:]
	if url == "/gedcom.js" {
		return &module, nil
	} else {
		f, err := files.Open(url)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		return javascript.ParseModule(parser.NewReaderTokeniser(f))
	}
}

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

type gedcomData []javascript.AssignmentExpression

func (g *gedcomData) Set(id uint64, data javascript.AssignmentExpression) {
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

func run() error {
	var (
		m      interface{} = &module
		err    error
		input  = flag.String("i", "-", "gedcom file")
		nowrap = flag.Bool("n", false, "just output gedcom data module, do not package all scripts")
		output = flag.String("o", "-", "output js file")
		f      = os.Stdin
		w      = os.Stdout
	)
	flag.Parse()
	if !*nowrap {
		if m, err = jslib.Package(jslib.Loader(opener), jslib.NoExports, jslib.File("/ged2web.js")); err != nil {
			return err
		}
	}
	if *input != "-" {
		if f, err = os.Open(*input); err != nil {
			return err
		}
	}
	err = processGedcom(f)
	f.Close()
	if err != nil {
		return err
	}
	if *output != "-" {
		if w, err = os.Create(*output); err != nil {
			return err
		}
	}
	fmt.Fprintf(w, "%s", m)
	return nil
}

func processGedcom(f io.Reader) error {
	indiIDs := make(idMap)
	famIDs := make(idMap)
	noneStr := javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenStr("")})}
	noneNum := javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(0)})}
	indis := gedcomData{{ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{ElementList: []javascript.AssignmentExpression{noneStr, noneStr, noneStr, noneStr, noneNum, noneNum}})}}
	fams := gedcomData{{ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{ElementList: []javascript.AssignmentExpression{noneNum, noneNum}})}}
	r := gedcom.NewReader(f, gedcom.AllowMissingRequired, gedcom.IgnoreInvalidValue, gedcom.AllowUnknownCharset, gedcom.AllowTerminatorsInValue, gedcom.AllowWrongLength, gedcom.AllowInvalidEscape, gedcom.AllowInvalidChars)
	for {
		record, err := r.Record()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		switch t := record.(type) {
		case *gedcom.Individual:
			person := append(make([]javascript.AssignmentExpression, 0, 6+len(t.SpouseOf)), noneStr, noneStr, noneStr, noneStr, noneNum, noneNum)
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
				person[0].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenStr(firstName)})
				person[1].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenStr(lastName)})
			}
			if t.Death.Date != "" {
				person[2].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenStr(strings.TrimSpace(string(t.Birth.Date)))})
				person[3].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenStr((string(t.Death.Date)))})
			}
			gender := uint64(1)
			switch t.Gender {
			case "F", "f", "Female", "FEMALE", "female":
				gender = 2
				fallthrough
			case "M", "m", "Male", "MALE", "male":
				person[4].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(gender)})
			}
			if len(t.ChildOf) > 0 {
				person[5].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(famIDs.GetID(t.ChildOf[0].ID))})
			}
			for _, spouse := range t.SpouseOf {
				person = append(person, javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(famIDs.GetID(spouse.ID))})})
			}
			indis.Set(indiIDs.GetID(t.ID), javascript.AssignmentExpression{
				ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{
					ElementList: person,
				}),
			})
		case *gedcom.Family:
			family := append(make([]javascript.AssignmentExpression, 0, 2+len(t.Children)), noneNum, noneNum)
			if t.Husband != "" {
				family[0].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(indiIDs.GetID(t.Husband))})
			}
			if t.Wife != "" {
				family[1].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(indiIDs.GetID(t.Wife))})
			}
			for _, child := range t.Children {
				family = append(family, javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: tokenNum(indiIDs.GetID(child))})})
			}
			fams.Set(famIDs.GetID(t.ID), javascript.AssignmentExpression{
				ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{
					ElementList: family,
				}),
			})
		}
	}
	ia.ElementList = indis
	fa.ElementList = fams
	return nil
}

func token(data string) *javascript.Token {
	return &javascript.Token{Token: parser.Token{Data: data}}
}

func tokenStr(data string) *javascript.Token {
	return token(strconv.Quote(data))
}

func tokenNum(data uint64) *javascript.Token {
	return token(strconv.FormatUint(data, 10))
}
