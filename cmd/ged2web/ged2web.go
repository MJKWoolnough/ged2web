package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"vimagination.zapto.org/gedcom"
	"vimagination.zapto.org/javascript"
	"vimagination.zapto.org/parser"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type idMap map[gedcom.Xref]string

func (i idMap) GetID(ref gedcom.Xref) string {
	id, ok := i[ref]
	if !ok {
		id = strconv.FormatUint(uint64(len(i))+1, 10)
		i[ref] = id
	}
	return id
}

func run() error {
	f, err := os.Open(os.Args[1])
	if err != nil {
		return err
	}
	m, err := makeModule(f)
	f.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%s", m)
	return nil
}

func makeModule(f io.Reader) (*javascript.Module, error) {
	indiIDs := make(idMap)
	famIDs := make(idMap)
	noneStr := javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token("\"\"")})}
	noneNum := javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token("0")})}
	indis := []javascript.AssignmentExpression{{ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{ElementList: []javascript.AssignmentExpression{noneStr, noneStr, noneStr, noneStr, noneNum, noneNum}})}}
	fams := []javascript.AssignmentExpression{{ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{ElementList: []javascript.AssignmentExpression{noneNum, noneNum}})}}
	r := gedcom.NewReader(f, gedcom.AllowMissingRequired, gedcom.IgnoreInvalidValue, gedcom.AllowUnknownCharset, gedcom.AllowTerminatorsInValue, gedcom.AllowWrongLength, gedcom.AllowInvalidEscape, gedcom.AllowInvalidChars)
	for {
		record, err := r.Record()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		switch t := record.(type) {
		case *gedcom.Individual:
			indiIDs.GetID(t.ID)
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
				person[0].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(strconv.Quote(firstName))})
				person[1].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(strconv.Quote(lastName))})
			}
			if t.Death.Date != "" {
				person[2].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(strconv.Quote(strings.TrimSpace(string(t.Birth.Date))))})
				person[3].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(strconv.Quote(strings.TrimSpace(string(t.Death.Date))))})
			}
			gender := "1"
			switch t.Gender {
			case "F", "f", "Female", "FEMALE", "female":
				gender = "2"
				fallthrough
			case "M", "m", "Male", "MALE", "male":
				person[4].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(gender)})
			}
			if len(t.ChildOf) > 0 {
				person[5].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(famIDs.GetID(t.ChildOf[0].ID))})
			}
			for _, spouse := range t.SpouseOf {
				person = append(person, javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(famIDs.GetID(spouse.ID))})})
			}
			indis = append(indis, javascript.AssignmentExpression{
				ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{
					ElementList: person,
				}),
			})
		case *gedcom.Family:
			famIDs.GetID(t.ID)
			family := append(make([]javascript.AssignmentExpression, 0, 2+len(t.Children)), noneNum, noneNum)
			if t.Husband != "" {
				family[0].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(indiIDs.GetID(t.Husband))})
			}
			if t.Wife != "" {
				family[1].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(indiIDs.GetID(t.Wife))})
			}
			for _, child := range t.Children {
				family = append(family, javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(indiIDs.GetID(child))})})
			}
			fams = append(fams, javascript.AssignmentExpression{
				ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{
					ElementList: family,
				}),
			})
		}
	}
	return &javascript.Module{
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
										ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{ElementList: indis}),
									},
								},
								{
									BindingIdentifier: token("families"),
									Initializer: &javascript.AssignmentExpression{
										ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{ElementList: fams}),
									},
								},
							},
						},
					},
				},
			},
		},
	}, nil
}

func token(data string) *javascript.Token {
	return &javascript.Token{Token: parser.Token{Data: data}}
}
