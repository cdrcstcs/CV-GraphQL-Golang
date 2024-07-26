package database

import (
	"context"
	"log"
	"time"

	"github.com/akhil/gql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb+srv://username:password@cluster0.yourcluster.mongodb.net/school-management"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

// Student CRUD Operations
func (db *DB) GetStudent(id string) *model.Student {
	studentCollec := db.client.Database("school-management").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var student model.Student
	err := studentCollec.FindOne(ctx, filter).Decode(&student)
	if err != nil {
		log.Fatal(err)
	}
	return &student
}

func (db *DB) GetStudents() []*model.Student {
	studentCollec := db.client.Database("school-management").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var students []*model.Student
	cursor, err := studentCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &students); err != nil {
		log.Fatal(err)
	}

	return students
}

func (db *DB) CreateStudent(studentInfo model.CreateStudentInput) *model.Student {
	studentCollec := db.client.Database("school-management").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := studentCollec.InsertOne(ctx, bson.M{
		"firstName":        studentInfo.FirstName,
		"lastName":         studentInfo.LastName,
		"email":            studentInfo.Email,
		"birthdate":        studentInfo.Birthdate,
		"enrollmentStatus": studentInfo.EnrollmentStatus,
		"guardians":        studentInfo.Guardians,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnStudent := model.Student{
		ID:               insertedID,
		FirstName:        studentInfo.FirstName,
		LastName:         studentInfo.LastName,
		Email:            studentInfo.Email,
		Birthdate:        studentInfo.Birthdate,
		EnrollmentStatus: studentInfo.EnrollmentStatus,
		Guardians:        studentInfo.Guardians,
	}
	return &returnStudent
}

func (db *DB) UpdateStudent(studentId string, studentInfo model.UpdateStudentInput) *model.Student {
	studentCollec := db.client.Database("school-management").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateStudentInfo := bson.M{}

	if studentInfo.FirstName != nil {
		updateStudentInfo["firstName"] = studentInfo.FirstName
	}
	if studentInfo.LastName != nil {
		updateStudentInfo["lastName"] = studentInfo.LastName
	}
	if studentInfo.Email != nil {
		updateStudentInfo["email"] = studentInfo.Email
	}
	if studentInfo.Birthdate != nil {
		updateStudentInfo["birthdate"] = studentInfo.Birthdate
	}
	if studentInfo.EnrollmentStatus != nil {
		updateStudentInfo["enrollmentStatus"] = studentInfo.EnrollmentStatus
	}
	if studentInfo.Guardians != nil {
		updateStudentInfo["guardians"] = studentInfo.Guardians
	}

	_id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateStudentInfo}

	results := studentCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var student model.Student

	if err := results.Decode(&student); err != nil {
		log.Fatal(err)
	}

	return &student
}

func (db *DB) DeleteStudent(studentId string) *model.DeleteStudentResponse {
	studentCollec := db.client.Database("school-management").Collection("students")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(studentId)
	filter := bson.M{"_id": _id}
	_, err := studentCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteStudentResponse{DeletedStudentID: studentId}
}

// Teacher CRUD Operations
func (db *DB) GetTeacher(id string) *model.Teacher {
	teacherCollec := db.client.Database("school-management").Collection("teachers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var teacher model.Teacher
	err := teacherCollec.FindOne(ctx, filter).Decode(&teacher)
	if err != nil {
		log.Fatal(err)
	}
	return &teacher
}

func (db *DB) GetTeachers() []*model.Teacher {
	teacherCollec := db.client.Database("school-management").Collection("teachers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var teachers []*model.Teacher
	cursor, err := teacherCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &teachers); err != nil {
		log.Fatal(err)
	}

	return teachers
}

func (db *DB) CreateTeacher(teacherInfo model.CreateTeacherInput) *model.Teacher {
	teacherCollec := db.client.Database("school-management").Collection("teachers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := teacherCollec.InsertOne(ctx, bson.M{
		"firstName":  teacherInfo.FirstName,
		"lastName":   teacherInfo.LastName,
		"email":      teacherInfo.Email,
		"officeHours": teacherInfo.OfficeHours,
		"subjects":   teacherInfo.Subjects,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnTeacher := model.Teacher{
		ID:          insertedID,
		FirstName:   teacherInfo.FirstName,
		LastName:    teacherInfo.LastName,
		Email:       teacherInfo.Email,
		OfficeHours: teacherInfo.OfficeHours,
		Subjects:    teacherInfo.Subjects,
	}
	return &returnTeacher
}

func (db *DB) UpdateTeacher(teacherId string, teacherInfo model.UpdateTeacherInput) *model.Teacher {
	teacherCollec := db.client.Database("school-management").Collection("teachers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateTeacherInfo := bson.M{}

	if teacherInfo.FirstName != nil {
		updateTeacherInfo["firstName"] = teacherInfo.FirstName
	}
	if teacherInfo.LastName != nil {
		updateTeacherInfo["lastName"] = teacherInfo.LastName
	}
	if teacherInfo.Email != nil {
		updateTeacherInfo["email"] = teacherInfo.Email
	}
	if teacherInfo.OfficeHours != nil {
		updateTeacherInfo["officeHours"] = teacherInfo.OfficeHours
	}
	if teacherInfo.Subjects != nil {
		updateTeacherInfo["subjects"] = teacherInfo.Subjects
	}

	_id, _ := primitive.ObjectIDFromHex(teacherId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateTeacherInfo}

	results := teacherCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var teacher model.Teacher

	if err := results.Decode(&teacher); err != nil {
		log.Fatal(err)
	}

	return &teacher
}

func (db *DB) DeleteTeacher(teacherId string) *model.DeleteTeacherResponse {
	teacherCollec := db.client.Database("school-management").Collection("teachers")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(teacherId)
	filter := bson.M{"_id": _id}
	_, err := teacherCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteTeacherResponse{DeletedTeacherID: teacherId}
}

// Course CRUD Operations
func (db *DB) GetCourse(id string) *model.Course {
	courseCollec := db.client.Database("school-management").Collection("courses")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var course model.Course
	err := courseCollec.FindOne(ctx, filter).Decode(&course)
	if err != nil {
		log.Fatal(err)
	}
	return &course
}

func (db *DB) GetCourses() []*model.Course {
	courseCollec := db.client.Database("school-management").Collection("courses")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var courses []*model.Course
	cursor, err := courseCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &courses); err != nil {
		log.Fatal(err)
	}

	return courses
}

func (db *DB) CreateCourse(courseInfo model.CreateCourseInput) *model.Course {
	courseCollec := db.client.Database("school-management").Collection("courses")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := courseCollec.InsertOne(ctx, bson.M{
		"title":        courseInfo.Title,
		"description":  courseInfo.Description,
		"teacherId":    courseInfo.TeacherId,
		"prerequisites": courseInfo.Prerequisites,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnCourse := model.Course{
		ID:            insertedID,
		Title:         courseInfo.Title,
		Description:   courseInfo.Description,
		TeacherId:     courseInfo.TeacherId,
		Prerequisites: courseInfo.Prerequisites,
	}
	return &returnCourse
}

func (db *DB) UpdateCourse(courseId string, courseInfo model.UpdateCourseInput) *model.Course {
	courseCollec := db.client.Database("school-management").Collection("courses")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateCourseInfo := bson.M{}

	if courseInfo.Title != nil {
		updateCourseInfo["title"] = courseInfo.Title
	}
	if courseInfo.Description != nil {
		updateCourseInfo["description"] = courseInfo.Description
	}
	if courseInfo.TeacherId != nil {
		updateCourseInfo["teacherId"] = courseInfo.TeacherId
	}
	if courseInfo.Prerequisites != nil {
		updateCourseInfo["prerequisites"] = courseInfo.Prerequisites
	}

	_id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateCourseInfo}

	results := courseCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var course model.Course

	if err := results.Decode(&course); err != nil {
		log.Fatal(err)
	}

	return &course
}

func (db *DB) DeleteCourse(courseId string) *model.DeleteCourseResponse {
	courseCollec := db.client.Database("school-management").Collection("courses")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(courseId)
	filter := bson.M{"_id": _id}
	_, err := courseCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteCourseResponse{DeletedCourseID: courseId}
}

// Class CRUD Operations
func (db *DB) GetClass(id string) *model.Class {
	classCollec := db.client.Database("school-management").Collection("classes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var class model.Class
	err := classCollec.FindOne(ctx, filter).Decode(&class)
	if err != nil {
		log.Fatal(err)
	}
	return &class
}

func (db *DB) GetClasses() []*model.Class {
	classCollec := db.client.Database("school-management").Collection("classes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var classes []*model.Class
	cursor, err := classCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &classes); err != nil {
		log.Fatal(err)
	}

	return classes
}

func (db *DB) CreateClass(classInfo model.CreateClassInput) *model.Class {
	classCollec := db.client.Database("school-management").Collection("classes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := classCollec.InsertOne(ctx, bson.M{
		"courseId":  classInfo.CourseId,
		"schedule":  classInfo.Schedule,
		"location":  classInfo.Location,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnClass := model.Class{
		ID:        insertedID,
		CourseId:  classInfo.CourseId,
		Schedule:  classInfo.Schedule,
		Location:  classInfo.Location,
	}
	return &returnClass
}

func (db *DB) UpdateClass(classId string, classInfo model.UpdateClassInput) *model.Class {
	classCollec := db.client.Database("school-management").Collection("classes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateClassInfo := bson.M{}

	if classInfo.Schedule != nil {
		updateClassInfo["schedule"] = classInfo.Schedule
	}
	if classInfo.Location != nil {
		updateClassInfo["location"] = classInfo.Location
	}

	_id, _ := primitive.ObjectIDFromHex(classId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateClassInfo}

	results := classCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var class model.Class

	if err := results.Decode(&class); err != nil {
		log.Fatal(err)
	}

	return &class
}

func (db *DB) DeleteClass(classId string) *model.DeleteClassResponse {
	classCollec := db.client.Database("school-management").Collection("classes")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(classId)
	filter := bson.M{"_id": _id}
	_, err := classCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteClassResponse{DeletedClassID: classId}
}

// Grade CRUD Operations
func (db *DB) GetGrade(id string) *model.Grade {
	gradeCollec := db.client.Database("school-management").Collection("grades")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var grade model.Grade
	err := gradeCollec.FindOne(ctx, filter).Decode(&grade)
	if err != nil {
		log.Fatal(err)
	}
	return &grade
}

func (db *DB) GetGrades() []*model.Grade {
	gradeCollec := db.client.Database("school-management").Collection("grades")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var grades []*model.Grade
	cursor, err := gradeCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &grades); err != nil {
		log.Fatal(err)
	}

	return grades
}

func (db *DB) CreateGrade(gradeInfo model.CreateGradeInput) *model.Grade {
	gradeCollec := db.client.Database("school-management").Collection("grades")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := gradeCollec.InsertOne(ctx, bson.M{
		"courseId":   gradeInfo.CourseId,
		"studentId":  gradeInfo.StudentId,
		"score":      gradeInfo.Score,
		"gradeLetter": gradeInfo.GradeLetter,
		"comments":   gradeInfo.Comments,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnGrade := model.Grade{
		ID:          insertedID,
		CourseId:    gradeInfo.CourseId,
		StudentId:   gradeInfo.StudentId,
		Score:       gradeInfo.Score,
		GradeLetter: gradeInfo.GradeLetter,
		Comments:    gradeInfo.Comments,
	}
	return &returnGrade
}

func (db *DB) UpdateGrade(gradeId string, gradeInfo model.UpdateGradeInput) *model.Grade {
	gradeCollec := db.client.Database("school-management").Collection("grades")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateGradeInfo := bson.M{}

	if gradeInfo.Score != nil {
		updateGradeInfo["score"] = gradeInfo.Score
	}
	if gradeInfo.GradeLetter != nil {
		updateGradeInfo["gradeLetter"] = gradeInfo.GradeLetter
	}
	if gradeInfo.Comments != nil {
		updateGradeInfo["comments"] = gradeInfo.Comments
	}

	_id, _ := primitive.ObjectIDFromHex(gradeId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateGradeInfo}

	results := gradeCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var grade model.Grade

	if err := results.Decode(&grade); err != nil {
		log.Fatal(err)
	}

	return &grade
}

func (db *DB) DeleteGrade(gradeId string) *model.DeleteGradeResponse {
	gradeCollec := db.client.Database("school-management").Collection("grades")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(gradeId)
	filter := bson.M{"_id": _id}
	_, err := gradeCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteGradeResponse{DeletedGradeID: gradeId}
}

// Guardian CRUD Operations
func (db *DB) GetGuardian(id string) *model.Guardian {
	guardianCollec := db.client.Database("school-management").Collection("guardians")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id}
	var guardian model.Guardian
	err := guardianCollec.FindOne(ctx, filter).Decode(&guardian)
	if err != nil {
		log.Fatal(err)
	}
	return &guardian
}

func (db *DB) GetGuardians() []*model.Guardian {
	guardianCollec := db.client.Database("school-management").Collection("guardians")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var guardians []*model.Guardian
	cursor, err := guardianCollec.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	if err = cursor.All(ctx, &guardians); err != nil {
		log.Fatal(err)
	}

	return guardians
}

func (db *DB) CreateGuardian(guardianInfo model.CreateGuardianInput) *model.Guardian {
	guardianCollec := db.client.Database("school-management").Collection("guardians")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserg, err := guardianCollec.InsertOne(ctx, bson.M{
		"firstName": guardianInfo.FirstName,
		"lastName":  guardianInfo.LastName,
		"email":     guardianInfo.Email,
		"phone":     guardianInfo.Phone,
		"relationship": guardianInfo.Relationship,
	})

	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserg.InsertedID.(primitive.ObjectID).Hex()
	returnGuardian := model.Guardian{
		ID:             insertedID,
		FirstName:      guardianInfo.FirstName,
		LastName:       guardianInfo.LastName,
		Email:          guardianInfo.Email,
		Phone:          guardianInfo.Phone,
		Relationship:   guardianInfo.Relationship,
	}
	return &returnGuardian
}

func (db *DB) UpdateGuardian(guardianId string, guardianInfo model.UpdateGuardianInput) *model.Guardian {
	guardianCollec := db.client.Database("school-management").Collection("guardians")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateGuardianInfo := bson.M{}

	if guardianInfo.FirstName != nil {
		updateGuardianInfo["firstName"] = guardianInfo.FirstName
	}
	if guardianInfo.LastName != nil {
		updateGuardianInfo["lastName"] = guardianInfo.LastName
	}
	if guardianInfo.Email != nil {
		updateGuardianInfo["email"] = guardianInfo.Email
	}
	if guardianInfo.Phone != nil {
		updateGuardianInfo["phone"] = guardianInfo.Phone
	}
	if guardianInfo.Relationship != nil {
		updateGuardianInfo["relationship"] = guardianInfo.Relationship
	}

	_id, _ := primitive.ObjectIDFromHex(guardianId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateGuardianInfo}

	results := guardianCollec.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var guardian model.Guardian

	if err := results.Decode(&guardian); err != nil {
		log.Fatal(err)
	}

	return &guardian
}

func (db *DB) DeleteGuardian(guardianId string) *model.DeleteGuardianResponse {
	guardianCollec := db.client.Database("school-management").Collection("guardians")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(guardianId)
	filter := bson.M{"_id": _id}
	_, err := guardianCollec.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteGuardianResponse{DeletedGuardianID: guardianId}
}
