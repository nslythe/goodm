package goodm

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_GetColletionName_1(t *testing.T) {
	type Test struct{}
	if GetCollectionName(Obj(&Test{})) != "goodm__test" {
		t.Error("failed")
	}
}

func Test_GetColletionName_2(t *testing.T) {
	type Test_allo struct{}
	if GetCollectionName(Obj(&Test_allo{})) != "goodm__test_allo" {
		t.Error("failed")
	}
}

func Test_GetColletionName_3(t *testing.T) {
	type tests struct{}
	if GetCollectionName(Obj(&tests{})) != "goodm_tests" {
		t.Error("failed")
	}
}

func Test_GetColletionName_4(t *testing.T) {
	type TestTest struct{}
	if GetCollectionName(Obj(&TestTest{})) != "goodm__test_test" {
		t.Error("failed")
	}
}

func Test_GetColletionName_5(t *testing.T) {
	type TestTest1 struct{}
	if GetCollectionName(Obj(&TestTest1{})) != "goodm__test_test1" {
		t.Error("failed")
	}
}

func Test_GetColletionName_6(t *testing.T) {
	type TestTest1 struct{}
	if GetCollectionName(Obj(&TestTest1{})) != "goodm__test_test1" {
		t.Error("failed")
	}
	if GetCollectionName(Obj(&[]TestTest1{})) != "goodm__test_test1" {
		t.Error("failed")
	}
}

func Test_save_1(t *testing.T) {
	type TestTest1 struct{}
	obj := Obj(&TestTest1{})

	err := Coll(obj).Save(obj)
	if err == nil {
		t.Error("No BaseObject in parent")
	}
}

func Test_save_2(t *testing.T) {
	type TestTest1 struct {
		BaseObject
	}
	obj := Obj(&TestTest1{})
	err := Coll(obj).Save(obj)
	if err == nil {
		t.Errorf("BaseObject not inline %s", err)
	}
}

func Test_save_21(t *testing.T) {
	type TestTest1 struct {
		BaseObject
	}
	obj := &TestTest1{}
	err := Coll(obj).Save(obj)
	if err == nil {
		t.Errorf("BaseObject not inline %s", err)
	}
}

func Test_save_22(t *testing.T) {
	type TestTest1 struct {
		BaseObject `goodm-collection:"test_tesfdfsdfsdt11"`
	}
	if GetCollectionName(Obj(TestTest1{})) != "test_tesfdfsdfsdt11" {
		t.Error()
	}
}

func Test_save_3(t *testing.T) {
	type TestTest1 struct {
		BaseObject `bson:"inline"`
	}

	obj := Obj(&TestTest1{})
	err := Coll(obj).Save(obj)
	if err != nil {
		t.Errorf("BaseObject not inline %s", err)
	}
}

func Test_save_31(t *testing.T) {
	type TestTest1 struct {
		BaseObject `bson:"inline" goodm-collection:"inheritance_test_test1"`
		Test1      string
	}
	type TestTest2 struct {
		TestTest1 `bson:"inline"`
		Test2     string
	}

	t1 := TestTest1{
		Test1: "Test1",
	}
	t2 := TestTest2{
		Test2: "Test2",
		TestTest1: TestTest1{
			Test1: "Test21",
		},
	}

	obj1 := Obj(&t1)
	obj2 := Obj(&t2)

	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Errorf("BaseObject not inline %s", err)
	}
	err = Coll(obj2).Save(obj2)
	if err != nil {
		t.Errorf("BaseObject not inline %s", err)
	}

	t_test1 := TestTest1{}
	t_test1.BaseObject.Id = t1.Id

	Coll(t_test1).Load(&t_test1)
	if t_test1.Test1 != "Test1" {
		t.Error()
	}

	t_test1.BaseObject.Id = t2.Id

	Coll(t_test1).Load(&t_test1)
	if t_test1.Test1 != "Test21" {
		t.Error()
	}
}

func Test_save_33(t *testing.T) {
	type TestTest1 struct {
		BaseObject `bson:"inline" goodm-collection:"inheritance_test_test33"`
		Test1      string
	}
	type TestTest2 struct {
		BaseObject `bson:"inline" goodm-collection:"inheritance_test_test33"`
		Test1      string
		Test2      string
	}

	t1 := TestTest1{
		Test1: "Test1",
	}
	t2 := TestTest2{
		Test1: "Test1",
		Test2: "Test2",
	}
	obj1 := Obj(&t1)
	obj2 := Obj(&t2)

	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Errorf("BaseObject not inline %s", err)
	}
	err = Coll(obj2).Save(obj2)
	if err != nil {
		t.Errorf("BaseObject not inline %s", err)
	}

	t1.Test1 = ""
	t2.Test2 = ""

	err = Coll(obj1).FindSpecificType(obj1, primitive.M{"test1": "Test1"})
	if err != nil {
		t.Error(err)
	}
	if t1.Test1 != "Test1" {
		t.Error()
	}

	Coll(obj2).FindSpecificType(obj2, primitive.M{"test1": "Test1"})
	if t2.Test2 != "Test2" {
		t.Error()
	}

}

func Test_save_4(t *testing.T) {
	type TestTest1 struct {
		BaseObject `bson:"inline"`
		TestStr    string
	}
	test_obj := TestTest1{}

	obj1 := Obj(&test_obj)
	obj1.Field("TestStr").Set("TestStr")
	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Errorf("BaseObject not inline %s", err)
	}

	obj2 := Obj(&TestTest1{})
	err = Coll(obj2).Load(obj2)
	if err == nil {
		t.Error()
	}

	id, err := obj1.GetID()
	obj2.SetID(id)
	err = Coll(obj2).Load(obj2)
	if err != nil {
		t.Error()
	}

	s1 := obj1.Field("TestStr").Interface().(string)
	s2 := obj2.Field("TestStr").Interface().(string)
	if s1 != s2 {
		t.Error()
	}
}

func Test_save_5(t *testing.T) {
	type TestTest_save_5 struct {
		BaseObject `bson:"inline"`
		TestStr    string
	}

	test := TestTest_save_5{}
	test_slice := []TestTest_save_5{}

	obj := Obj(&test)
	obj_slice := Obj(&test_slice)

	Coll(obj).Drop()

	test.TestStr = "TestStr"
	err := Coll(obj).Save(obj)
	if err != nil {
		t.Error(err)
	}

	err = Coll(obj_slice).Find(obj_slice, primitive.M{})
	if err != nil {
		t.Error(err)
	}
	if len(test_slice) != 1 {
		t.Error()
	}
}

func Test_save_6(t *testing.T) {
	type TestTest_save_5 struct {
		BaseObject `bson:"inline"`
		TestStr    string
	}

	test := TestTest_save_5{}
	obj := Obj(&test)

	Coll(obj).Drop()

	test.TestStr = "TestStr"
	err := Coll(obj).Save(obj)
	if err != nil {
		t.Error(err)
	}

	var test2 TestTest_save_5
	obj2 := Obj(&test2)
	err = Coll(obj2).Find(obj2, primitive.M{})
	if err != nil {
		t.Error(err)
	}
	if test2.TestStr != "TestStr" {
		t.Error()
	}
}

func Test_save_7(t *testing.T) {
	type Test_save_7 struct {
		BaseObject `bson:"inline"`
		TestStr    string
	}

	test1 := Test_save_7{}
	test1.TestStr = "TestStr"
	obj1 := Obj(&test1)

	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Error(err)
	}

	test_slice := []Test_save_7{}
	test_slice = append(test_slice, Test_save_7{})
	obj_slice := Obj(&test_slice)
	err = Coll(obj_slice).Find(obj_slice, primitive.M{})
	if err != nil {
		t.Error(err)
	}
	l := len(test_slice)
	if l != 1 {
		t.Error()
	}
}

func Test_save_8(t *testing.T) {
	type Test_save_8 struct {
		BaseObject `bson:"inline"`
		TestStr    string
	}

	test := Test_save_8{}
	obj1 := Obj(&test)
	Coll(obj1).Drop()

	test2 := Test_save_8{}
	test2.TestStr = "TestStr"
	obj2 := Obj(&test2)
	err := Coll(obj2).Save(obj2)
	if err != nil {
		t.Error(err)
	}

	err = Coll(obj1).Delete(obj2)

	test_slice := []Test_save_8{}
	obj_slice := Obj(&test_slice)
	err = Coll(obj_slice).Find(obj_slice, primitive.M{})
	if err != nil {
		t.Error(err)
	}
	l := len(test_slice)
	if l != 0 {
		t.Error()
	}
}

func Test_save_9(t *testing.T) {
	type Test_save_9 struct {
		BaseObject `bson:"inline"`
		TestStr    string
	}
	test1 := Test_save_9{}
	test1.TestStr = "TestStr1"
	obj1 := Obj(&test1)

	test2 := Test_save_9{}
	test2.TestStr = "TestStr2"
	obj2 := Obj(&test2)

	test3 := Test_save_9{}
	test3.TestStr = "TestStr3"
	obj3 := Obj(&test3)

	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Error(err)
	}

	err = Coll(obj2).Save(obj2)
	if err != nil {
		t.Error(err)
	}

	err = Coll(obj3).Save(obj3)
	if err != nil {
		t.Error(err)
	}

	test_find := []Test_save_9{}
	obj_find := Obj(&test_find)
	err = Coll(obj_find).Find(obj_find, primitive.M{})
	l := len(test_find)
	if l != 3 {
		t.Error()
	}

	obj_to_delete := []Test_save_9{}
	obj_to_delete = append(obj_to_delete, Test_save_9{
		BaseObject: BaseObject{
			Id: test1.Id,
		},
	})
	obj_to_delete = append(obj_to_delete, Test_save_9{
		BaseObject: BaseObject{
			Id: test2.Id,
		},
	})
	obj_to_delete = append(obj_to_delete, Test_save_9{
		BaseObject: BaseObject{
			Id: test3.Id,
		},
	})

	err = Coll(obj1).Delete(Obj(&obj_to_delete))

	err = Coll(obj_find).Find(obj_find, primitive.M{})
	if err != nil {
		t.Error(err)
	}
	l = len(test_find)
	if l != 0 {
		t.Error()
	}
}

func Test_update_1(t *testing.T) {
	type Test_update_1 struct {
		BaseObject `bson:"inline"`
		TestStr    string
		Num        int
	}

	t1 := Test_update_1{
		TestStr: "1",
		Num:     0,
	}
	t2 := Test_update_1{
		TestStr: "2",
		Num:     2,
	}
	obj1 := Obj(&t1)
	obj2 := Obj(&t2)
	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Error(err)
	}

	err = Coll(obj2).Update(obj2, primitive.M{"teststr": "1"})
	if err != nil {
		t.Error(err)
	}
	if t2.BaseObject.Id != t1.BaseObject.Id {
		t.Error()
	}

	err = Coll(obj1).Update(obj1, primitive.M{"teststr": "2"})
	if err != nil {
		t.Error(err)
	}
	if t2.BaseObject.Id != t1.BaseObject.Id {
		t.Error()
	}
}

func Test_count_1(t *testing.T) {
	type Test_count_1 struct {
		BaseObject `bson:"inline"`
		TestStr    string
		Num        int
	}

	t1 := Test_count_1{
		TestStr: "1",
		Num:     0,
	}
	t2 := Test_count_1{
		TestStr: "2",
		Num:     2,
	}
	obj1 := Obj(&t1)
	obj2 := Obj(&t2)
	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Error(err)
	}
	err = Coll(obj2).Save(obj2)
	if err != nil {
		t.Error(err)
	}

	if c, _ := Coll(obj1).Count(primitive.M{}); c != 2 {
		t.Error()
	}
}

func Test_update_all_1(t *testing.T) {
	type Test_update_all_1 struct {
		BaseObject `bson:"inline"`
		TestStr    string
		Num        int
	}

	t1 := Test_update_all_1{
		TestStr: "1",
		Num:     0,
	}
	t2 := Test_update_all_1{
		TestStr: "1",
		Num:     2,
	}
	obj1 := Obj(&t1)
	obj2 := Obj(&t2)
	Coll(obj1).Drop()

	err := Coll(obj1).Save(obj1)
	if err != nil {
		t.Error(err)
	}
	err = Coll(obj2).Save(obj2)
	if err != nil {
		t.Error(err)
	}

	t3 := Test_update_all_1{
		TestStr: "1",
		Num:     3,
	}

	n, err := Coll(obj2).UpdateAll(&t3, primitive.M{"teststr": "1"})
	if err != nil {
		t.Error(err)
	}
	if n != 2 {
		t.Error()
	}

	err = Coll(obj2).Load(obj2)
	if err != nil {
		t.Error(err)
	}
	if t2.Num != 3 {
		t.Error()
	}

	err = Coll(obj1).Load(obj1)
	if err != nil {
		t.Error(err)
	}
	if t1.Num != 3 {
		t.Error()
	}
}

func Test_coll_string_1(t *testing.T) {
	type Test_coll_string_1 struct {
		BaseObject `bson:"inline"`
		TestStr    string
		Num        int
	}

	t1 := Test_coll_string_1{
		TestStr: "1",
		Num:     0,
	}
	err := Coll("test").Save(&t1)
	if err != nil {
		t.Error(err)
	}
}
