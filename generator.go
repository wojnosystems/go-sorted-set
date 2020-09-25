//go:generate go-poor-generics g --package sorted_set --outFile sorted.go --templateFile template.txt --headerFile header-template.txt --namesToPrimitiveTypes "Int=int,Int64=int64,Int32=int32,Int16=int16,Int8=int8,Uint=uint,Uint64=uint64,Uint32=uint32,Uint16=uint16,Uint8=uint8,Byte=byte,Rune=rune,String=string,Float64=float64,Float32=float32,Uintptr=uintptr"
package sorted_set
