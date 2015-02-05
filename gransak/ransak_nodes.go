package filter

type OperatorFunction func(re *RansakEmulator)

type Node struct {
	Name       string
	Nodes      []*Node
	Apply      OperatorFunction
	IsOperator bool
}

var Tree = &Node{
	Name: "Operators",
	Nodes: []*Node{
		&Node{
			Name: "or",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.template += "OR "
			},
		},
		&Node{
			Name: "and",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.template += "AND "
			},
		},
		&Node{
			Name: "eq",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "in",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("IN (" + re.placeholder + ")")
			},
		},
		&Node{
			Name: "matches",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.placeholder + "'")
			},
		},
		&Node{
			Name: "cont",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.placeholder + "%'")
			},
			IsOperator: true,
			Nodes: []*Node{
				&Node{
					Name: "any",
					Apply: func(re *RansakEmulator) {
						field := re.getLastField()

						values := re.param.parts
						times := len(values) - 1

						statement := field + " LIKE '%" + re.placeholder + "%'"

						re.template += statement

						for i := 0; i < (times); i++ {
							re.template += " OR " + statement
						}
					},
				},
			},
		},
		&Node{
			Name: "lt",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("< " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "lteq",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("<= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "gt",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("> " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "gteq",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder(">= " + re.getCorrectSqlFormat(re.placeholder))
			},
		},
		&Node{
			Name: "start",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '" + re.placeholder + "%'")
			},
		},
		&Node{
			Name: "end",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("LIKE '%" + re.placeholder + "'")
			},
		},
		&Node{
			Name: "true",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("= 't'")
			},
		},
		&Node{
			Name: "false",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("= 'f'")
			},
		},
		&Node{
			Name: "present",
			Apply: func(re *RansakEmulator) {
				field := re.appendField()
				re.replacePlaceholder("IS NOT NULL AND " + re.placeholder)
				re.replacePlaceholder(field + " <> ''")
			},
		},
		&Node{
			Name: "blank",
			Apply: func(re *RansakEmulator) {
				field := re.appendField()
				re.replacePlaceholder("IS NULL OR " + re.placeholder)
				re.replacePlaceholder(field + " = ''")
			},
		},
		&Node{
			Name: "null",
			Apply: func(re *RansakEmulator) {
				re.appendField()
				re.replacePlaceholder("IS NULL")
			},
		},
		&Node{
			Name: "not",
			Nodes: []*Node{
				&Node{
					Name: "eq",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("<> " + re.getCorrectSqlFormat(re.placeholder))
					},
				},
				&Node{
					Name: "in",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("NOT IN (" + re.placeholder + ")")
					},
				},
				&Node{
					Name: "cont",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '%" + re.placeholder + "%'")
					},
					IsOperator: true,
					Nodes: []*Node{
						&Node{
							Name: "any",
							Apply: func(re *RansakEmulator) {
								field := re.getLastField()

								values := re.param.parts
								times := len(values) - 1

								statement := field + " NOT LIKE '%" + re.placeholder + "%'"

								re.template += statement

								for i := 0; i < (times); i++ {
									re.template += " AND " + statement
								}
							},
						},
					},
				},
				&Node{
					Name: "start",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '" + re.placeholder + "%'")
					},
				},
				&Node{
					Name: "end",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("NOT LIKE '%" + re.placeholder + "'")
					},
				},
				&Node{
					Name: "true",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("<> 't'")
					},
				},
				&Node{
					Name: "false",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("<> 'f'")
					},
				},
				&Node{
					Name: "null",
					Apply: func(re *RansakEmulator) {
						re.appendField()
						re.replacePlaceholder("IS NOT NULL")
					},
				},
			},
		},
		&Node{
			Name: "does",
			Nodes: []*Node{
				&Node{
					Name: "not",
					Nodes: []*Node{
						&Node{
							Name: "match",
							Apply: func(re *RansakEmulator) {
								re.appendField()
								re.replacePlaceholder("NOT LIKE '" + re.placeholder + "'")
							},
						},
					},
				},
			},
		},
	},
}
