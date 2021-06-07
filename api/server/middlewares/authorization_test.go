package middlewares

import (
	"net/http"
	"reflect"
	"spotifo/types"
	"testing"
)

func TestAuthorization_fetchUserInfo(t *testing.T) {
	type fields struct {
		next     http.Handler
		endpoint string
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    types.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			fields: fields{
				next:     nil,
				endpoint: "https://patient-seal-72.hasura.app/api/rest/user",
			},
			args: args{
				token: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6InlEdkIxWnNvSWhkYzVmT2xEM3dEZyJ9.eyJodHRwczovL2hhc3VyYS5pby9qd3QvY2xhaW1zIjp7IngtaGFzdXJhLWRlZmF1bHQtcm9sZSI6InVzZXIiLCJ4LWhhc3VyYS1hbGxvd2VkLXJvbGVzIjpbInVzZXIiXSwieC1oYXN1cmEtdXNlci1pZCI6Imdvb2dsZS1vYXV0aDJ8MTEwMjI0NjcyNTU5MTE3Njk4MjAzIn0sImlzcyI6Imh0dHBzOi8vaGlyb3N1bWUuanAuYXV0aDAuY29tLyIsInN1YiI6Imdvb2dsZS1vYXV0aDJ8MTEwMjI0NjcyNTU5MTE3Njk4MjAzIiwiYXVkIjpbImh0dHBzOi8vaGFzdXJhLmlvL2xlYXJuIiwiaHR0cHM6Ly9oaXJvc3VtZS5qcC5hdXRoMC5jb20vdXNlcmluZm8iXSwiaWF0IjoxNjIyOTk3NTU2LCJleHAiOjE2MjMwMDQ3NTYsImF6cCI6ImplbEhDZkdmZk1SZVduT3F0TEQ3WlU4em1lVW5LUWNsIiwic2NvcGUiOiJvcGVuaWQgZW1haWwifQ.IHJA1E4zbHdO1GsJev2cQQyJEeGQG3j85KQNcrP5aoIXQd1d74ieajCfzisSgDzy9rOh7Va7AX92_bpmIG1dXx9qgafMhkEo95es4lzl9mgkihhmQWx39bXtoeD8_roGqbO9xTE4dGT2_FyNVIw2KUg0XbWXWRzI41h8dMZSY8vZXxMBm3A6Etxw6VA_ra5rSCSdaPBZZG1-nCUhw1gB_8i7DKpRPR3FkwVdWAnxIhMUziHK_SY0kQ-GdZNru62Gil8gbe1HVnhreOZQ2d5qvmm7cWlEBI6nSjIURP5cLhxZAEr5t3r_2kRXDpiLrkrvE_Q6lLWnEwfG3VF4lMsLtQ",
			},
			want: types.User{
				Email:     "bdqa1030@gmail.com",
				FullName:  "Cường Vũ",
				Id:        2,
				CreatedAt: "2021-06-06T16:36:59.443315",
				Avatar:    "https://lh3.googleusercontent.com/a/AATXAJxb9Zjwe_zDnatQbDpaBQp92lgrBPPQwpCY4V1l=s96-c",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Authorization{
				Next:     tt.fields.next,
				Endpoint: tt.fields.endpoint,
			}
			got, err := a.fetchUserInfo(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
