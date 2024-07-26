package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/akhil/gql/database"
	"github.com/akhil/gql/graph/generated"
	"github.com/akhil/gql/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*model.DeleteUserResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// CreateStudent is the resolver for the createStudent field.
func (r *mutationResolver) CreateStudent(ctx context.Context, input model.CreateStudentInput) (*model.Student, error) {
	panic(fmt.Errorf("not implemented: CreateStudent - createStudent"))
}

// UpdateStudent is the resolver for the updateStudent field.
func (r *mutationResolver) UpdateStudent(ctx context.Context, id string, input model.UpdateStudentInput) (*model.Student, error) {
	panic(fmt.Errorf("not implemented: UpdateStudent - updateStudent"))
}

// DeleteStudent is the resolver for the deleteStudent field.
func (r *mutationResolver) DeleteStudent(ctx context.Context, id string) (*model.DeleteStudentResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteStudent - deleteStudent"))
}

// CreateTeacher is the resolver for the createTeacher field.
func (r *mutationResolver) CreateTeacher(ctx context.Context, input model.CreateTeacherInput) (*model.Teacher, error) {
	panic(fmt.Errorf("not implemented: CreateTeacher - createTeacher"))
}

// UpdateTeacher is the resolver for the updateTeacher field.
func (r *mutationResolver) UpdateTeacher(ctx context.Context, id string, input model.UpdateTeacherInput) (*model.Teacher, error) {
	panic(fmt.Errorf("not implemented: UpdateTeacher - updateTeacher"))
}

// DeleteTeacher is the resolver for the deleteTeacher field.
func (r *mutationResolver) DeleteTeacher(ctx context.Context, id string) (*model.DeleteTeacherResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteTeacher - deleteTeacher"))
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.CreateCourseInput) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: CreateCourse - createCourse"))
}

// UpdateCourse is the resolver for the updateCourse field.
func (r *mutationResolver) UpdateCourse(ctx context.Context, id string, input model.UpdateCourseInput) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: UpdateCourse - updateCourse"))
}

// DeleteCourse is the resolver for the deleteCourse field.
func (r *mutationResolver) DeleteCourse(ctx context.Context, id string) (*model.DeleteCourseResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteCourse - deleteCourse"))
}

// CreateClass is the resolver for the createClass field.
func (r *mutationResolver) CreateClass(ctx context.Context, input model.CreateClassInput) (*model.Class, error) {
	panic(fmt.Errorf("not implemented: CreateClass - createClass"))
}

// UpdateClass is the resolver for the updateClass field.
func (r *mutationResolver) UpdateClass(ctx context.Context, id string, input model.UpdateClassInput) (*model.Class, error) {
	panic(fmt.Errorf("not implemented: UpdateClass - updateClass"))
}

// DeleteClass is the resolver for the deleteClass field.
func (r *mutationResolver) DeleteClass(ctx context.Context, id string) (*model.DeleteClassResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteClass - deleteClass"))
}

// CreateGrade is the resolver for the createGrade field.
func (r *mutationResolver) CreateGrade(ctx context.Context, input model.CreateGradeInput) (*model.Grade, error) {
	panic(fmt.Errorf("not implemented: CreateGrade - createGrade"))
}

// UpdateGrade is the resolver for the updateGrade field.
func (r *mutationResolver) UpdateGrade(ctx context.Context, id string, input model.UpdateGradeInput) (*model.Grade, error) {
	panic(fmt.Errorf("not implemented: UpdateGrade - updateGrade"))
}

// DeleteGrade is the resolver for the deleteGrade field.
func (r *mutationResolver) DeleteGrade(ctx context.Context, id string) (*model.DeleteGradeResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteGrade - deleteGrade"))
}

// CreateGuardian is the resolver for the createGuardian field.
func (r *mutationResolver) CreateGuardian(ctx context.Context, input model.CreateGuardianInput) (*model.Guardian, error) {
	panic(fmt.Errorf("not implemented: CreateGuardian - createGuardian"))
}

// UpdateGuardian is the resolver for the updateGuardian field.
func (r *mutationResolver) UpdateGuardian(ctx context.Context, id string, input model.UpdateGuardianInput) (*model.Guardian, error) {
	panic(fmt.Errorf("not implemented: UpdateGuardian - updateGuardian"))
}

// DeleteGuardian is the resolver for the deleteGuardian field.
func (r *mutationResolver) DeleteGuardian(ctx context.Context, id string) (*model.DeleteGuardianResponse, error) {
	panic(fmt.Errorf("not implemented: DeleteGuardian - deleteGuardian"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - user"))
}

// Students is the resolver for the students field.
func (r *queryResolver) Students(ctx context.Context) ([]*model.Student, error) {
	panic(fmt.Errorf("not implemented: Students - students"))
}

// Student is the resolver for the student field.
func (r *queryResolver) Student(ctx context.Context, id string) (*model.Student, error) {
	panic(fmt.Errorf("not implemented: Student - student"))
}

// Teachers is the resolver for the teachers field.
func (r *queryResolver) Teachers(ctx context.Context) ([]*model.Teacher, error) {
	panic(fmt.Errorf("not implemented: Teachers - teachers"))
}

// Teacher is the resolver for the teacher field.
func (r *queryResolver) Teacher(ctx context.Context, id string) (*model.Teacher, error) {
	panic(fmt.Errorf("not implemented: Teacher - teacher"))
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	panic(fmt.Errorf("not implemented: Courses - courses"))
}

// Course is the resolver for the course field.
func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	panic(fmt.Errorf("not implemented: Course - course"))
}

// Classes is the resolver for the classes field.
func (r *queryResolver) Classes(ctx context.Context) ([]*model.Class, error) {
	panic(fmt.Errorf("not implemented: Classes - classes"))
}

// Class is the resolver for the class field.
func (r *queryResolver) Class(ctx context.Context, id string) (*model.Class, error) {
	panic(fmt.Errorf("not implemented: Class - class"))
}

// Grades is the resolver for the grades field.
func (r *queryResolver) Grades(ctx context.Context) ([]*model.Grade, error) {
	panic(fmt.Errorf("not implemented: Grades - grades"))
}

// Grade is the resolver for the grade field.
func (r *queryResolver) Grade(ctx context.Context, id string) (*model.Grade, error) {
	panic(fmt.Errorf("not implemented: Grade - grade"))
}

// Guardians is the resolver for the guardians field.
func (r *queryResolver) Guardians(ctx context.Context) ([]*model.Guardian, error) {
	panic(fmt.Errorf("not implemented: Guardians - guardians"))
}

// Guardian is the resolver for the guardian field.
func (r *queryResolver) Guardian(ctx context.Context, id string) (*model.Guardian, error) {
	panic(fmt.Errorf("not implemented: Guardian - guardian"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
var db = database.Connect()

func (r *mutationResolver) CreateJobListing(ctx context.Context, input model.CreateJobListingInput) (*model.JobListing, error) {
	return db.CreateJobListing(input), nil
}
func (r *mutationResolver) UpdateJobListing(ctx context.Context, id string, input model.UpdateJobListingInput) (*model.JobListing, error) {
	return db.UpdateJobListing(id, input), nil
}
func (r *mutationResolver) DeleteJobListing(ctx context.Context, id string) (*model.DeleteJobResponse, error) {
	return db.DeleteJobListing(id), nil
}
func (r *queryResolver) Jobs(ctx context.Context) ([]*model.JobListing, error) {
	return db.GetJobs(), nil
}
func (r *queryResolver) Job(ctx context.Context, id string) (*model.JobListing, error) {
	return db.GetJob(id), nil
}
