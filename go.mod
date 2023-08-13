module my-go-app/example 

require (
   myFunc.com/ufunc v0.0.0
) 

replace (
   myFunc.com/ufunc v0.0.0 => ./ufunc
)

go 1.20
