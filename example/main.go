package main

import (
	"fmt"
	"github.com/PucklaMotzer09/GLSLGenerator"
)

func main() {
	var vertex glslgen.VertexGenerator
	var fragment glslgen.FragmentGenerator

	vertex.SetVersion(
		"110",
	).AddAttributes(
		[]glslgen.Variable{
			glslgen.Variable{"vec3", "highp", "vertex"},
			glslgen.Variable{"vec3", "highp", "normal"},
			glslgen.Variable{"vec2", "highp", "texCoord"},
			glslgen.Variable{"vec3", "highp", "tangent"},
		},
	).AddModule(
		glslgen.Module{
			[]glslgen.Variable{
				glslgen.Variable{"mat4", "highp", "transformMatrix3D"},
				glslgen.Variable{"mat4", "highp", "viewMatrix3D"},
				glslgen.Variable{"mat4", "highp", "projectionMatrix3D"},
			},
			[]glslgen.Function{
				glslgen.Function{
					"void doStuff()",
					"gl_Position = projectionMatrix3D*viewMatrix3D*transformMatrix3D*vec4(vertex,1.0);",
				},
			},
			"passthrough",
			"doStuff();",
		},
	)

	fragment.SetVersion(
		"110",
	).AddModule(
		glslgen.Module{
			[]glslgen.Variable{},
			[]glslgen.Function{},
			"passthrough",
			"gl_FragColor = vec4(1.0,0.0,0.0,1.0);",
		},
	)

	fmt.Print("---------\n", vertex.String(), "-------------\n\n")
	fmt.Print("---------\n", fragment.String(), "-------------\n\n")
}
