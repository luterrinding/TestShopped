package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"github.com/graphql-go/graphql"
)

type App struct {
	scenariosList []scenariosList
	PriceList []PriceList
  }

type PriceList struct {
	SKU      string
	Name    string
	Price   float64
	Qty 	int
}

type Chart struct {
	SKU      string
	Qty 	int
	Shema     string
}

type scenariosList struct {
	SKU    string
	Skufreeproduct  string
	Maxbuy   int
	Status 	string
	Discountpersent int
}

type responsePrice struct {
	Data struct {
		List []PriceList `json:"list"`
	} `json:"data"`
}



type responseScenario struct {
	Data struct {
		List []scenariosList `json:"list"`
	} `json:"data"`
}

var priceType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PriceList",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"qty": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)


var scenarioType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "scenariosList",
		Fields: graphql.Fields{
			"sku": &graphql.Field{
				Type: graphql.String,
			},
			"skufreeproduct": &graphql.Field{
				Type: graphql.String,
			},
			"maxbuy": &graphql.Field{
				Type: graphql.Int,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"discountpersent": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)


func populatedata() []PriceList {
	prices1 := PriceList{
		SKU:     "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty: 10,
	}
	prices2 := PriceList{ 
		SKU:     "43N23P",
		Name:  "MacBook Pro",
		Price: 5399.99,
		Qty: 5,
	}
	prices3 := PriceList{
		SKU:     "A304SD",
		Name:  "Alexa Speaker",
		Price: 109.50,
		Qty: 10,
	}
	prices4 := PriceList{
		SKU:     "234234",
		Name:  "Raspberry Pi B",
		Price: 30.00,
		Qty: 2,
	}

	var pricelists []PriceList
	pricelists = append(pricelists, prices1)
	pricelists = append(pricelists, prices2)
	pricelists = append(pricelists, prices3)
	pricelists = append(pricelists, prices4)
	return pricelists
}

func scenariosData() []scenariosList {
	case1 := scenariosList{
		SKU:     "43N23P",
		Skufreeproduct:  "234234",
		Maxbuy: 1,
		Status: "free_item",
		Discountpersent: 0,
	}

	case2 := scenariosList{
		SKU:     "120P90",
		Skufreeproduct:  "120P90",
		Maxbuy: 3,
		Status: "buy_item_free_item",
		Discountpersent: 0,
	}

	case3 := scenariosList{
		SKU:     "A304SD",
		Skufreeproduct:  "",
		Maxbuy: 3,
		Status: "item_discount",
		Discountpersent: 10,
	}
	

	var scenariolist []scenariosList
	scenariolist = append(scenariolist, case1)
	scenariolist = append(scenariolist, case2)
	scenariolist = append(scenariolist, case3)
	return scenariolist
}

func Pembulatan2(x float64) float64{
	return math.Round(x*100)/100
}


func (a App) GetPriceItem(itemsku string)  float64 {
	for _, row := range a.PriceList {
		if itemsku == row.SKU {
			return row.Price
		}
	}
	return 0
}

func RoundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}


func (a App) Scenarios1(itemCart map[string]int)  float64 {
	var totalPrice float64
	var totalBonus float64
	for item, qty := range itemCart {
		priceI := a.GetPriceItem(item)
		totalPrice += Pembulatan2(priceI * float64(qty))
		for _, row := range a.scenariosList {
			if item == row.SKU {
				//fmt.Println("aaa  \n", row.Status, row.Maxbuy )

				// Scenarios 1
				if row.Status == "free_item" &&  row.Maxbuy >= qty {
					bonus := a.GetPriceItem(row.Skufreeproduct) 
					qItemBonus := qty / row.Maxbuy
					totalBonus += bonus * float64(qItemBonus)
				}


				// Scenarios 2
				if row.Status == "buy_item_free_item" &&  row.Maxbuy == qty {
					bonus := a.GetPriceItem(row.Skufreeproduct) 
					qItemBonus := qty / row.Maxbuy
					totalBonus +=Pembulatan2((bonus * float64(qItemBonus)))
				}

				// Scenarios 3
				if row.Status == "item_discount" &&  row.Maxbuy == qty {
					bonus := a.GetPriceItem(row.SKU)
					dis := Pembulatan2(float64(row.Discountpersent) / 100 )
					totalBonus +=  dis * ( bonus * float64(qty))
				}
				break
			}
		}
	}
	return RoundUp((totalPrice - totalBonus),2)
}

func main() {
	app := &App{}
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
      fmt.Println(err)
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
      fmt.Println(err)
    }
	app.scenariosList = dataModel.Data.List

	cartItem := map[string]int{"43N23P": 1 , "234234": 1}
	totalbayar := app.Scenarios1(cartItem)
	fmt.Println("scenario 1", totalbayar)

	cartItem = map[string]int{"120P90": 3 }
	totalbayar  = app.Scenarios1(cartItem)
	fmt.Println("scenario 2", totalbayar)
	
	cartItem = map[string]int{"A304SD": 3 }
	totalbayar = app.Scenarios1(cartItem)
	fmt.Println("scenario 3", totalbayar)

}

