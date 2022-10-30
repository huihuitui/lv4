package main

type intList []string

func (m intList) Len() int {
	return len(m)
}
func (m intList) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m intList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]

}
