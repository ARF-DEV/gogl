package gogl

import (
	"errors"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

func CompileShader(shaderSource string, shaderType uint32) (uint32, error) {
	shaderId := gl.CreateShader(shaderType)
	shaderSource += "\x00"
	cSource, free := gl.Strs(shaderSource)
	gl.ShaderSource(shaderId, 1, cSource, nil)
	free()
	gl.CompileShader(shaderId)
	var glErr int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &glErr)
	if glErr == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)
		log := string(make([]byte, logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, nil, gl.Str(log))
		return 0, errors.New("Error Failed to compile shader:\n" + log)
	}
	return shaderId, nil
}
