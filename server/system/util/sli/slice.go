package sli

import (
	"errors"
	"log"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"chunDoc/system/util/str"
)

//删除
func DelInt(slice []int, sep int) []int {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == sep {
		return []int{}
	}
	var n_slice = []int{}
	for i := range slice {
		if slice[i] == sep && i == count {
			return slice[:count]
		} else if slice[i] == sep {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return n_slice
}

func JoinInt(s []int, sep ...string) string {
	l := ","
	if len(sep) > 0 {
		l = sep[0]
	}
	str := make([]string, 0)
	for _, v := range s {
		str = append(str, strconv.FormatInt(int64(v), 10))
	}
	return strings.Join(str, l)
}

func SplitInt(str string, sep ...string) ([]int, error) {
	l := ","
	if len(sep) > 0 {
		l = sep[0]
	}
	i := make([]int, 0)
	s := strings.Split(str, l)
	for _, v := range s {
		d, err := strconv.Atoi(v)
		if err != nil {
			return []int{}, err
		}
		i = append(i, d)
	}
	return i, nil
}

func InSlice(v interface{}, s interface{}) bool {
	val := reflect.Indirect(reflect.ValueOf(v))

	if val.Kind() == reflect.Slice {
		log.Println("Warning: first parameter cannot be a slice, Function: InSlice")
		return false
	}

	sv := reflect.Indirect(reflect.ValueOf(s))
	if sv.Kind() != reflect.Slice {
		log.Println("Warning: second parameter needs a slice, Function: InSlice")
		return false
	}

	st := reflect.TypeOf(s).Elem().String()
	vt := reflect.TypeOf(v).String()
	if st != vt {
		log.Println("Warning: two parameters are not the same type, Function: InSlice")
		return false
	}

	switch vt {
	case "string":
		for _, vv := range s.([]string) {
			if vv == v {
				return true
			}
		}
	case "int":
		for _, vv := range s.([]int) {
			if vv == v {
				return true
			}
		}
	case "int64":
		for _, vv := range s.([]int64) {
			if vv == v {
				return true
			}
		}
	case "float64":
		for _, vv := range s.([]float64) {
			if vv == v {
				return true
			}
		}
	default:
		log.Println("Warning: this type is not supported, Function: InSlice")
		return false
	}

	return false
}

func InInterface(v interface{}, sl []interface{}) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func Slice(iface interface{}) ([]interface{}, error) {
	d := make([]interface{}, 0)
	v := reflect.Indirect(reflect.ValueOf(iface))
	if v.Kind() != reflect.Slice {
		return []interface{}{}, errors.New("Needs a slice")
	}
	t := reflect.TypeOf(iface).Elem().String()
	switch t {
	case "string":
		for _, v := range iface.([]string) {
			if strings.Trim(str.String(v), " ") == "" {
				continue
			}
			d = append(d, v)
		}
	case "int":
		for _, v := range iface.([]int) {
			d = append(d, v)
		}
	case "int64":
		for _, v := range iface.([]int64) {
			d = append(d, v)
		}
	case "float64":
		for _, v := range iface.([]float64) {
			d = append(d, v)
		}
	default:
		return d, errors.New("this type is not supported")
	}
	return d, nil
}

func Unique(list *[]string) []string {
	r := make([]string, 0)
	temp := map[string]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

func UniqueInt(list *[]int) []int {
	r := make([]int, 0)
	temp := map[int]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

/*func UniqueInt(list *[]int) []int {
	var r []int = []int{}
	for _, i := range *list {
		if len(r) == 0 {
			r = append(r, i)
		} else {
			for k, v := range r {
				if i == v {
					break
				}
				if k == len(r)-1 {
					r = append(r, i)
				}
			}
		}
	}
	return r
}*/

func UniqueInt64(list *[]int64) []int64 {
	r := make([]int64, 0)
	temp := map[int64]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

func UniqueIface(list *[]interface{}) []interface{} {
	r := make([]interface{}, 0)
	temp := map[interface{}]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

func UniqueFloat(list *[]float64) []float64 {
	r := make([]float64, 0)
	temp := map[float64]byte{}
	for _, v := range *list {
		l := len(temp)
		temp[v] = 0
		if len(temp) != l {
			r = append(r, v)
		}
	}
	return r
}

func Chunk(slice []interface{}, size int) (chunk [][]interface{}) {
	if size <= 0 {
		chunk = make([][]interface{}, 0)
		return
	}
	if size >= len(slice) {
		chunk = append(chunk, slice)
		return
	}
	end := size
	for i := 0; i < len(slice); i += size {
		if end > len(slice) {
			chunk = append(chunk, slice[i:])
			break
		}
		chunk = append(chunk, slice[i:end])
		end += size
	}
	return
}

func Rand(a []string) (b string) {
	r := rand.Intn(len(a))
	b = a[r]
	return
}

func RandInt(a []int) (b int) {
	r := rand.Intn(len(a))
	b = a[r]
	return
}

func RandInt64(a []int64) (b int64) {
	r := rand.Intn(len(a))
	b = a[r]
	return
}

func RandList(min, max int) []int {
	if max < min {
		min, max = max, min
	}
	length := max - min + 1
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	list := rand.Perm(length)
	for index := range list {
		list[index] += min
	}
	return list
}

func Sum(i []int64) (s int64) {
	for _, v := range i {
		s += v
	}
	return
}

func Diff(slice1, slice2 []string) (diffslice []string) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func DiffInt(slice1, slice2 []int) (diffslice []int) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func DiffInt64(slice1, slice2 []int64) (diffslice []int64) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func DiffFloat(slice1, slice2 []float64) (diffslice []float64) {
	for _, v := range slice1 {
		if !InSlice(v, slice2) {
			diffslice = append(diffslice, v)
		}
	}
	return
}

func Intersect(slice1, slice2 []string) (interSlice []string) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			interSlice = append(interSlice, v)
		}
	}
	return
}

func IntersectInt(slice1, slice2 []int) (interSlice []int) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			interSlice = append(interSlice, v)
		}
	}
	return
}

func IntersectIn64(slice1, slice2 []int64) (interSlice []int64) {
	for _, v := range slice1 {
		if InSlice(v, slice2) {
			interSlice = append(interSlice, v)
		}
	}
	return
}

func Reverse(s *[]string) []string {
	l := len(*s) - 1
	m := len(*s) / 2
	r := make([]string, len(*s))
	for i := 0; i < l; i++ {
		j := l - i
		if i == j || i >= m {
			break
		}
		r[i], r[j] = (*s)[j], (*s)[i]
	}
	return r
}

func ReverseInt(s *[]int) []int {
	l := len(*s) - 1
	m := len(*s) / 2
	r := make([]int, len(*s))
	for i := 0; i < l; i++ {
		j := l - i
		if i == j || i >= m {
			break
		}
		r[i], r[j] = (*s)[j], (*s)[i]
	}
	return r
}

func ReverseInt64(s *[]int64) []int64 {
	l := len(*s) - 1
	m := len(*s) / 2
	r := make([]int64, len(*s))
	for i := 0; i < l; i++ {
		j := l - i
		if i == j || i >= m {
			break
		}
		r[i], r[j] = (*s)[j], (*s)[i]
	}
	return r
}

func ReverseFloat(s *[]float64) []float64 {
	l := len(*s) - 1
	m := len(*s) / 2
	r := make([]float64, len(*s))
	for i := 0; i < l; i++ {
		j := l - i
		if i == j || i >= m {
			break
		}
		r[i], r[j] = (*s)[j], (*s)[i]
	}
	return r
}

func Range(start, end, step int64) (intslice []int64) {
	for i := start; i <= end; i += step {
		intslice = append(intslice, i)
	}
	return
}

func Pad(slice []interface{}, size int, val interface{}) []interface{} {
	if size <= len(slice) {
		return slice
	}
	for i := 0; i < (size - len(slice)); i++ {
		slice = append(slice, val)
	}
	return slice
}

func Shuffle(slice []interface{}) []interface{} {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
	return slice
}
