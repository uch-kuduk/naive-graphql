# naive-graphql
Naive graphql-api
## Начинаем работу (с "нуля" - репозитория не содержащего ничего на тему graphql):  
1)  ставим https://github.com/99designs/gqlgen
```bazaar
go get github.com/99designs/gqlgen  
```
2) Создаем файл gqlgen.yml  
3) 
```bazaar
go run github.com/99designs/gqlgen init
```
4) Куча ошибок, нестрашно, но уже сегенерировано:
   - graph/generated/generated.go  
   - graph/resolver.go  
   - graph/schema.graphqls   
5) 
```bazaar
go get -u ./...
```   
ТУТ уже можно запустить сервер go run server.go и зайти на playground  

6) Заполняем файлы как нам нужно:  
   - graph/scalar/email.go - пример собственного типа, упомянутого в gqlgen.yml и использованных в схемах  
   - graph/schema/directive.graphql - пример директив, используемых в схемах  
   - graph/schema/garage.graphql - пример схем типов  
   - graph/schema/schema.graphql - пример запросов  
7) Повторно запускаем генерацию:  
```bazaar
go run github.com/99designs/gqlgen generate  
```  
// ОШИБКИ //  
8) Выполнить все go get   
9) поменять в go.mod: 	github.com/vektah/gqlparser/v2 v2.1.0  // ==> not v2.2.0 !!!  
10) 
```bazaar
go run github.com/99designs/gqlgen generate
```  
11) убрать лишнее в sсhema_resolvers.go 
    
ТУТ снова можно запустить сервер go run server.go и зайти на playground и открыть вкладку schema  

14) Выполнить запрос:  
```bazaar
query carBrands {
    carBrands{brands{
      name}
    }
}
```
15) Имплементируем пару методов schema.resolvers.go: 
```go
func (r *queryResolver) CarBrands(ctx context.Context) (*model.CarBrandsPayload, error) {
	return &model.CarBrandsPayload{
		Brands: []*model.Brand{
			{"Lada", "Lada", "http://pngimg.com/uploads/lada/lada_PNG105.png", "lada@lada.ru"},
			{"Toyota","Toyota", "http://logocentral.info/wp-content/uploads/2018/12/Toyota-Logo-1989.jpg","toyota@mail.com"},
		},
	}, nil
}

func (r *queryResolver) UserCars(ctx context.Context) (*model.CarsPayload, error) {
	return nil, fmt.Errorf("not implemented")
}
```
16) выполнить запрос:  
```bazaar
query carBrands {
    carBrands{brands{
      name
    	email}
    }
}
