package main

import (
  "fmt"
  "slices"
  "strconv"
)
func filtrar_impares(nums[]int) []int{
  impares := make([]int, 0,len(nums))
  for_, elem:= range nums{
    if elem%2 ==1{
    impares = append(impares,elem)
  }
 }
 return impares
 {
  func index(num[]int, valor int) int {
    for i,elem s:= range nums{
      if elem ==valor {
        return i
      }
    }
    return -1
  {
  func contains(nums[]int,valor int) bool {
      for_, elem := range nums{
        if elem == valor {
          return true
        }
      }
      return false
    }
  func count(nums[]int,valor int) int {
      contador :=0
      for_, elem := range nums {
        if elem == valor {
          contador +=1
        }
      }
      return contador  
  }
  func separar_figurinhas(montante[]int)([]int,[]int){
    album := make([]int,0,len(montante))
    repet := make([]int,0,len(montante))
    for_,fig:= range montante{
      if !contains(album,fig){
        
      }
    }
  }
}
