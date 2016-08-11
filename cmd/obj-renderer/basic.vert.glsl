#version 400

uniform mat4 projectionMatrix;
uniform mat4 cameraMatrix;
uniform mat4 modelMatrix;
uniform mat3 normalMatrix;

in vec3 inPosition;
in vec3 inNormal;

flat out vec3 color;

void main() {
	gl_Position = projectionMatrix * cameraMatrix * modelMatrix * vec4(inPosition, 1);

	mat4 modelViewMatrix = modelMatrix * cameraMatrix;

	// Position of the vertex, in worldspace : M * position
	vec3 Position_worldspace = (modelMatrix * vec4(inPosition,1)).xyz;

	// Vector that goes from the vertex to the camera, in camera space.
	// In camera space, the camera is at the origin (0,0,0).
	vec3 vertexPosition_cameraspace = ( modelViewMatrix * vec4(inPosition ,1)).xyz;
	vec3 EyeDirection_cameraspace = vec3(0,0,0) - vertexPosition_cameraspace;

	vec3 LightPosition_worldspace = vec3(6, 3, 6);

	// Vector that goes from the vertex to the light, in camera space. M is ommited because it's identity.
	vec3 LightPosition_cameraspace = ( cameraMatrix * vec4(LightPosition_worldspace,1)).xyz;
	vec3 LightDirection_cameraspace = LightPosition_cameraspace + EyeDirection_cameraspace;

	// Normal of the the vertex, in camera space
	vec3 Normal_cameraspace = (modelViewMatrix * vec4(inNormal,0)).xyz; // Only correct if ModelMatrix does not scale the model ! Use its inverse transpose if not.

	 // Normal of the computed fragment, in camera space
	vec3 n = normalize( Normal_cameraspace );

	// Direction of the light (from the fragment to the light)
	vec3 l = normalize( LightDirection_cameraspace );

	vec3 E = normalize( EyeDirection_cameraspace);

	vec3 R = reflect(-l,n);

	float cosAlpha = clamp( dot( E,R ), 0,1 );

	float cosTheta = dot( n,l );

	vec3 lightColor = vec3(1.0, 0.0, 0.0);

	color = lightColor * cosTheta + lightColor * cosAlpha;
}
