package main

import (
	"testing"
	"log"
	"encoding/json"
	"github.com/graphql-go/graphql"
	)



var (
	app App  = App{}
    Scenarios1Seharusnya   float64 = 5399.99
    Scenarios2Seharusnya   float64 = 99.98
    Scenarios3Seharusnya   float64 = 295.65
)

func init() {
	
	pricelists := populatedata()
	fields := graphql.Fields{
		"item": &graphql.Field{
			Type:   priceType,
			Description: "Get Product By ID",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["sku"].(string)
				if ok {
					// Find tutorial
					for _, row := range pricelists {
						if row.SKU == id {
							return row, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(priceType),
			Description: "Get Product List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return pricelists, nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	query := `
		{
			list {
				sku
				name
				price
				qty
			}
		}
	`
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)
	dataModelItem := responsePrice{}
    err = json.Unmarshal(rJSON, &dataModelItem)
    if err != nil {
		log.Println(err)
    }
	app.PriceList = dataModelItem.Data.List

	////// ================End schema Item==========================
	scenarios := scenariosData()
	scenarioF := graphql.Fields{
		"scenario": &graphql.Field{
			Type:   scenarioType,
			Description: "Get Item By ID",
			Args: graphql.FieldConfigArgument{
				"sku": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id, ok := p.Args["sku"].(string)
				if ok {
					// Find tutorial
					for _, row := range scenarios {
						if row.SKU == id {
							return row, nil
						}
					}
				}
				return nil, nil
			},
		},
		"list": &graphql.Field{
			Type:        graphql.NewList(scenarioType),
			Description: "Get Tutorial List",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return scenarios, nil
			},
		},
	}

	
	rootQueryItem := graphql.ObjectConfig{Name: "RootQuery", Fields: scenarioF}
	schemaConfigItem := graphql.SchemaConfig{Query: graphql.NewObject(rootQueryItem)}
	schemaItem, err := graphql.NewSchema(schemaConfigItem)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	// Query
	querys := `
		{
			list {
				sku
				skufreeproduct
				maxbuy
				status
				discountpersent
			}
		}
	`
	paramsS := graphql.Params{Schema: schemaItem, RequestString: querys}
	r = graphql.Do(paramsS)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ = json.Marshal(r)
	dataModel := responseScenario{}
    err = json.Unmarshal(rJSON, &dataModel)
    if err != nil {
		log.Println(err)
    }
	app.scenariosList = dataModel.Data.List
}


func TestScenario1(t *testing.T) {
	cartItem := map[string]int{"43N23P": 1 , "234234": 1}
    t.Logf("Scenario1 : %.2f", app.Scenarios1(cartItem))
    if app.Scenarios1(cartItem) != Scenarios1Seharusnya {
        t.Errorf("SALAH! TestScenario1 harusnya %.2f", Scenarios1Seharusnya)
    }
	
}


func TestScenario2(t *testing.T) {
	cartItem := map[string]int{"120P90": 3 }
    t.Logf("Scenario2 : %.2f", app.Scenarios1(cartItem))
    if app.Scenarios1(cartItem) != Scenarios2Seharusnya {
        t.Errorf("SALAH! TestScenario2 harusnya %.2f", Scenarios2Seharusnya)
    }
	
}

func TestScenario3(t *testing.T) {
	cartItem := map[string]int{"A304SD": 3 }
    t.Logf("Scenario3 : %.2f", app.Scenarios1(cartItem))
    if app.Scenarios1(cartItem) != Scenarios3Seharusnya {
        t.Errorf("SALAH! TestScenario3 harusnya %.2f", Scenarios3Seharusnya)
    }
	
}