package main

import (
	"testing"
)

// func BenchmarkUpdateUserByValue(b *testing.B) {
// 	u := User{Name: "Alice", Age: 25}
// 	for i := 0; i < b.N; i++ {
// 		updateUserByValue(u)
// 	}
// }

// func BenchmarkUpdateUserByValueAndReturn(b *testing.B) {
// 	u := User{Name: "Alice", Age: 25}
// 	for i := 0; i < b.N; i++ {
// 		u = updateUserByValueAndReturn(u)
// 	}
// }

// func BenchmarkUpdateUserByPointer(b *testing.B) {
// 	u := User{Name: "Alice", Age: 25}
// 	for i := 0; i < b.N; i++ {
// 		updateUserByPointer(&u)
// 	}
// }

type SmallStruct struct {
	ID     int
	Name   string
	Price  float64
	Active bool
	Score  int
}

type MediumStruct struct {
	ID      int
	Name    string
	Price   float64
	Active  bool
	Tags    []string
	Details map[string]string
	Reviews []float64
	Age     int
	Email   string
	Phone   string
	Address string
	Rank    int
	Score   int
}

type LargeStruct struct {
	ID               int
	Name             string
	Price            float64
	Active           bool
	Tags             []string
	Details          map[string]string
	Reviews          []float64
	Address          string
	Phone            string
	Email            string
	Age              int
	Rank             int
	Score            int
	Friends          []string
	Hobbies          []string
	History          []string
	Settings         map[string]interface{}
	Data             []int
	IsValid          bool
	Location         string
	Department       string
	Position         string
	Skills           []string
	Projects         []string
	Languages        []string
	Birthdate        string
	Gender           string
	Nationality      string
	MaritalStatus    string
	Notes            string
	ExperienceYears  int
	Education        string
	Certifications   []string
	LastLogin        string
	IPAddress        string
	LastUpdated      string
	SocialMedia      []string
	Contacted        bool
	EmergencyContact string
	Permissions      []string
	Bio              string
	AvatarURL        string
	TagsMap          map[string]bool
}

func updateSmallStructByValue(s SmallStruct) SmallStruct {
	s.Name = "Updated"
	s.Score = 100
	return s
}

func updateSmallStructByPointer(s *SmallStruct) {
	s.Name = "Updated"
	s.Score = 100
}

func BenchmarkSmallStructPassByValue(b *testing.B) {
	s := SmallStruct{ID: 1, Name: "Alice", Price: 10.5, Active: true, Score: 50}
	for i := 0; i < b.N; i++ {
		s = updateSmallStructByValue(s)
	}
}

func BenchmarkSmallStructPassByPointer(b *testing.B) {
	s := &SmallStruct{ID: 1, Name: "Alice", Price: 10.5, Active: true, Score: 50}
	for i := 0; i < b.N; i++ {
		updateSmallStructByPointer(s)
	}
}

func updateMediumStructByValue(m MediumStruct) MediumStruct {
	m.Name = "Updated"
	m.Score = 100
	m.Tags = append(m.Tags, "new tag")
	return m
}

func updateMediumStructByPointer(m *MediumStruct) {
	m.Name = "Updated"
	m.Score = 100
	m.Tags = append(m.Tags, "new tag")
}

func BenchmarkMediumStructPassByValue(b *testing.B) {
	m := MediumStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		m = updateMediumStructByValue(m)
	}
}

func BenchmarkMediumStructPassByPointer(b *testing.B) {
	m := &MediumStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		updateMediumStructByPointer(m)
	}
}

func updateLargeStructByValue(l LargeStruct) LargeStruct {
	l.Name = "Updated"
	l.Score = 100
	l.Tags = append(l.Tags, "new tag")
	return l
}

func updateLargeStructByPointer(l *LargeStruct) {
	l.Name = "Updated"
	l.Score = 100
	l.Tags = append(l.Tags, "new tag")
}

func BenchmarkLargeStructPassByValue(b *testing.B) {
	l := LargeStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		l = updateLargeStructByValue(l)
	}
}

func BenchmarkLargeStructPassByPointer(b *testing.B) {
	l := &LargeStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		updateLargeStructByPointer(l)
	}
}

// The extra benchmarks.
func updateSmallStructReturnPointer(s SmallStruct) *SmallStruct {
	s.Name = "Updated"
	s.Score = 100
	return &s
}

func updateSmallStructReturnPointerByPointer(s *SmallStruct) *SmallStruct {
	s.Name = "Updated"
	s.Score = 100
	return s
}

func BenchmarkSmallStructPassByValueReturnPointer(b *testing.B) {
	s := SmallStruct{ID: 1, Name: "Alice", Price: 10.5, Active: true, Score: 50}
	for i := 0; i < b.N; i++ {
		_ = updateSmallStructReturnPointer(s)
	}
}

func BenchmarkSmallStructPassByPointerReturnPointer(b *testing.B) {
	s := &SmallStruct{ID: 1, Name: "Alice", Price: 10.5, Active: true, Score: 50}
	for i := 0; i < b.N; i++ {
		_ = updateSmallStructReturnPointerByPointer(s)
	}
}

func updateMediumStructReturnPointer(m MediumStruct) *MediumStruct {
	m.Name = "Updated"
	m.Score = 100
	m.Tags = append(m.Tags, "new tag")
	return &m
}

func updateMediumStructReturnPointerByPointer(m *MediumStruct) *MediumStruct {
	m.Name = "Updated"
	m.Score = 100
	m.Tags = append(m.Tags, "new tag")
	return m
}

func BenchmarkMediumStructPassByValueReturnPointer(b *testing.B) {
	m := MediumStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		_ = updateMediumStructReturnPointer(m)
	}
}

func BenchmarkMediumStructPassByPointerReturnPointer(b *testing.B) {
	m := &MediumStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		_ = updateMediumStructReturnPointerByPointer(m)
	}
}

func updateLargeStructReturnPointer(l LargeStruct) *LargeStruct {
	l.Name = "Updated"
	l.Score = 100
	l.Tags = append(l.Tags, "new tag")
	return &l
}

func updateLargeStructReturnPointerByPointer(l *LargeStruct) *LargeStruct {
	l.Name = "Updated"
	l.Score = 100
	l.Tags = append(l.Tags, "new tag")
	return l
}

func BenchmarkLargeStructPassByValueReturnPointer(b *testing.B) {
	l := LargeStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		_ = updateLargeStructReturnPointer(l)
	}
}

func BenchmarkLargeStructPassByPointerReturnPointer(b *testing.B) {
	l := &LargeStruct{
		ID: 1, Name: "Alice", Price: 10.5, Active: true, Tags: []string{"tag1", "tag2"}, Details: map[string]string{"key": "value"},
		Reviews: []float64{4.5, 5.0}, Age: 25, Email: "alice@example.com", Phone: "123-456-789", Address: "123 Main St", Rank: 1, Score: 50}
	for i := 0; i < b.N; i++ {
		_ = updateLargeStructReturnPointerByPointer(l)
	}
}
