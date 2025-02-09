package env

import (
	"os"
	"strconv"
)

// gets an env variable  has a string and receives a fallback
func GetString(key, fallback string) string{
  val, ok := os.LookupEnv(key)
  if !ok {
    return fallback
  }
  return val;
}


func GetInt(key string, fallback int) int {
  val, ok := os.LookupEnv(key)
  if !ok {
    return fallback
  }
  // gets a string and creates an integer or an error
  valAsInt, err := strconv.Atoi(val)

  if err != nil{
    return fallback
  }

  return valAsInt
}
