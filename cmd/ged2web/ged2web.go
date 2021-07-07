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
	indiIDs := map[gedcom.Xref]string{"": "0"}
	famIDs := map[gedcom.Xref]string{"": "0"}
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
			if _, ok := indiIDs[t.ID]; !ok {
				indiIDs[t.ID] = strconv.FormatUint(uint64(len(indiIDs)), 10)
			}
			person := make([]javascript.AssignmentExpression, 6, 6+len(t.SpouseOf))
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
			} else {
				person[0] = noneStr
				person[1] = noneStr
			}
			if t.Death.Date != "" {
				person[2].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(strconv.Quote(strings.TrimSpace(string(t.Birth.Date))))})
				person[3].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(strconv.Quote(strings.TrimSpace(string(t.Death.Date))))})
			} else {
				person[2] = noneStr
				person[3] = noneStr
			}
			var gender string
			switch t.Gender {
			case "M", "m", "Male", "MALE", "male":
				gender = "1"
			case "F", "f", "Female", "FEMALE", "female":
				gender = "2"
			default:
				gender = "0"
			}
			person[4].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(gender)})
			if len(t.ChildOf) > 0 {
				childOf, ok := famIDs[t.ChildOf[0].ID]
				if !ok {
					childOf = strconv.FormatUint(uint64(len(famIDs)), 10)
					famIDs[t.ChildOf[0].ID] = childOf
				}
				person[5].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(childOf)})
			} else {
				person[5] = noneNum
			}
			for _, spouse := range t.SpouseOf {
				spouseOf, ok := famIDs[spouse.ID]
				if !ok {
					spouseOf = strconv.FormatUint(uint64(len(famIDs)), 10)
					famIDs[spouse.ID] = spouseOf
				}
				person = append(person, javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(spouseOf)})})
			}
			indis = append(indis, javascript.AssignmentExpression{
				ConditionalExpression: javascript.WrapConditional(&javascript.ArrayLiteral{
					ElementList: person,
				}),
			})
		case *gedcom.Family:
			if _, ok := famIDs[t.ID]; !ok {
				famIDs[t.ID] = strconv.FormatUint(uint64(len(famIDs)), 10)
			}
			family := make([]javascript.AssignmentExpression, 2, 2+len(t.Children))
			if t.Husband != "" {
				husb, ok := indiIDs[t.Husband]
				if !ok {
					husb = strconv.FormatUint(uint64(len(indiIDs)), 10)
					indiIDs[t.Husband] = husb
				}
				family[0].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(husb)})
			} else {
				family[0] = noneNum
			}
			if t.Wife != "" {
				wife, ok := indiIDs[t.Wife]
				if !ok {
					wife = strconv.FormatUint(uint64(len(indiIDs)), 10)
					indiIDs[t.Wife] = wife
				}
				family[1].ConditionalExpression = javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(wife)})
			} else {
				family[1] = noneNum
			}
			for _, child := range t.Children {
				childID, ok := indiIDs[child]
				if !ok {
					childID = strconv.FormatUint(uint64(len(indiIDs)), 10)
					indiIDs[child] = childID
				}
				family = append(family, javascript.AssignmentExpression{ConditionalExpression: javascript.WrapConditional(&javascript.PrimaryExpression{Literal: token(childID)})})
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
