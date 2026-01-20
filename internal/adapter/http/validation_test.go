package http

import (
	"testing"

	"canteen-app/internal/adapter/http/common"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestLoginRequestValidation(t *testing.T) {
	tests := []struct {
		name           string
		data           common.LoginRequest
		wantErrorTag   string
		wantErrorField string
	}{
		{
			name: "success",
			data: common.LoginRequest{
				Login:    "slim",
				Password: "shadsdfy",
			},
		},

		{
			name: "requied login",
			data: common.LoginRequest{
				Login:    "",
				Password: "shadsdfy",
			},
			wantErrorTag:   "required",
			wantErrorField: "Login",
		},

		{
			name: "requied password",
			data: common.LoginRequest{
				Login:    "sdfsdf",
				Password: "",
			},
			wantErrorTag:   "required",
			wantErrorField: "Password",
		},

		{
			name: "max login len",
			data: common.LoginRequest{
				Login:    "LFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeOLFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeO",
				Password: "shadsdfy",
			},
			wantErrorTag:   "max",
			wantErrorField: "Login",
		},

		{
			name: "max password len",
			data: common.LoginRequest{
				Login:    "dsdfsd",
				Password: "LFW6uiS8dPUxlx1Q045bHhftolgjVveVyJ9GCso2fNO3aFxBCeOLFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeO",
			},
			wantErrorTag:   "max",
			wantErrorField: "Password",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val := NewValidator()

			err := val.Struct(tc.data)

			if tc.wantErrorTag == "" {
				assert.Nil(t, err)
			} else {
				validationErrors := err.(validator.ValidationErrors)
				assert.Equal(t, tc.wantErrorTag, validationErrors[0].Tag())
				assert.Equal(t, tc.wantErrorField, validationErrors[0].Field())
			}
		})
	}
}

func TestRegisterRequestValidation(t *testing.T) {
	tests := []struct {
		name           string
		data           common.RegisterRequest
		wantErrorTag   string
		wantErrorField string
	}{
		{
			name: "success",
			data: common.RegisterRequest{
				Login:    "slim",
				Password: "shadsdfy",
				Name:     "sdfsdf",
				Surname:  "sdfsdf",
				Role:     "admin",
			},
		},

		{
			name: "requied login",
			data: common.RegisterRequest{
				Login:    "",
				Password: "shadsdfy",
				Name:     "sdfsdf",
				Surname:  "sdfsdf",
				Role:     "admin",
			},
			wantErrorTag:   "required",
			wantErrorField: "Login",
		},

		{
			name: "requied password",
			data: common.RegisterRequest{
				Login:    "sdfsdf",
				Password: "",
				Name:     "sdfsdf",
				Surname:  "sdfsdf",
				Role:     "admin",
			},
			wantErrorTag:   "required",
			wantErrorField: "Password",
		},

		{
			name: "requied name",
			data: common.RegisterRequest{
				Login:    "sdfsdf",
				Password: "shadsdfy",
				Name:     "",
				Surname:  "sdfsdf",
				Role:     "admin",
			},
			wantErrorTag:   "required",
			wantErrorField: "Name",
		},

		{
			name: "requied surname",
			data: common.RegisterRequest{
				Login:    "sdfsdf",
				Password: "shadsdfy",
				Name:     "sdfsdf",
				Surname:  "",
				Role:     "admin",
			},
			wantErrorTag:   "required",
			wantErrorField: "Surname",
		},

		{
			name: "requied role",
			data: common.RegisterRequest{
				Login:    "sdfsdf",
				Password: "shadsdfy",
				Name:     "sdfsdf",
				Surname:  "sdsdf",
				Role:     "",
			},
			wantErrorTag:   "required",
			wantErrorField: "Role",
		},

		{
			name: "min login len",
			data: common.RegisterRequest{
				Login:    "s",
				Password: "shadsdfy",
				Name:     "sdfsdf",
				Surname:  "sdsdf",
				Role:     "admin",
			},
			wantErrorTag:   "min",
			wantErrorField: "Login",
		},

		{
			name: "max login len",
			data: common.RegisterRequest{
				Login:    "LFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeOLFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeO",
				Password: "shadsdfy",
				Name:     "sdfsdf",
				Surname:  "sdsdf",
				Role:     "admin",
			},
			wantErrorTag:   "max",
			wantErrorField: "Login",
		},

		{
			name: "min password len",
			data: common.RegisterRequest{
				Login:    "sfgdfg",
				Password: "sdfy",
				Name:     "sdfsdf",
				Surname:  "sdsdf",
				Role:     "admin",
			},
			wantErrorTag:   "min",
			wantErrorField: "Password",
		},

		{
			name: "max password len",
			data: common.RegisterRequest{
				Login:    "dsdfsd",
				Password: "LFW6uiS8dPUxlx1Q045bHhftolgjVveVyJ9GCso2fNO3aFxBCeOLFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeO",
				Name:     "sdfsdf",
				Surname:  "sdsdf",
				Role:     "admin",
			},
			wantErrorTag:   "max",
			wantErrorField: "Password",
		},

		{
			name: "max name len",
			data: common.RegisterRequest{
				Login:    "dsdfsd",
				Password: "sdfsdfsdf",
				Name:     "LFW6uiS8dPUxlx1Q045bHhftolgjVveVyJ9GCso2fNO3aFxBCeOLFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeO",
				Surname:  "sdsdf",
				Role:     "admin",
			},
			wantErrorTag:   "max",
			wantErrorField: "Name",
		},

		{
			name: "max surname len",
			data: common.RegisterRequest{
				Login:    "dsdfsd",
				Password: "sdfsdfsdf",
				Name:     "sdfdd",
				Surname:  "LFW6uiS8dPUxlx1Q045bHhftolgjVveVyJ9GCso2fNO3aFxBCeOLFW6uiS8dPUxlx1Q045bHhftolgVveVyJ9GCso2fNO3aFxBCeO",
				Role:     "admin",
			},
			wantErrorTag:   "max",
			wantErrorField: "Surname",
		},

		{
			name: "only alpha in name",
			data: common.RegisterRequest{
				Login:    "dsdfsd",
				Password: "sdfsdfsdf",
				Name:     "sdfdd2",
				Surname:  "ssdfdfdfs",
				Role:     "admin",
			},
			wantErrorTag:   "alpha",
			wantErrorField: "Name",
		},

		{
			name: "only alpha in surname",
			data: common.RegisterRequest{
				Login:    "dsdfsd",
				Password: "sdfsdfsdf",
				Name:     "sdfdd",
				Surname:  "ssdfdf3dfs",
				Role:     "admin",
			},
			wantErrorTag:   "alpha",
			wantErrorField: "Surname",
		},

		{
			name: "invalid role",
			data: common.RegisterRequest{
				Login:    "dsdfsd",
				Password: "sdfsdfsdf",
				Name:     "sdfdd",
				Surname:  "ssdfdfdfs",
				Role:     "sdfsdfsdf",
			},
			wantErrorTag:   "oneof",
			wantErrorField: "Role",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			val := NewValidator()

			err := val.Struct(tc.data)

			if tc.wantErrorTag == "" {
				assert.Nil(t, err)
			} else {
				validationErrors := err.(validator.ValidationErrors)
				assert.Equal(t, tc.wantErrorTag, validationErrors[0].Tag())
				assert.Equal(t, tc.wantErrorField, validationErrors[0].Field())
			}
		})
	}
}
