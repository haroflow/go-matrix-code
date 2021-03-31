// Adapted from: https://www.shadertoy.com/view/Xltfzj

#version 330

// Input vertex attributes (from vertex shader)
in vec2 fragTexCoord;
in vec4 fragColor;

// Input uniform values
uniform sampler2D texture0;
uniform vec4 colDiffuse;
uniform float blurSize;
uniform vec2 windowSize;

// Output fragment color
out vec4 finalColor;

// To brighten the final image
float multiplier = 8;

void main()
{
    float Pi = 6.28318530718; // Pi*2
    float Directions = 32.0; // BLUR DIRECTIONS (Default 16.0 - More is better but slower)
    float Quality = 8.0; // BLUR QUALITY (Default 4.0 - More is better but slower)
   
    vec2 Radius = blurSize / windowSize;

    // Pixel colour
    vec4 Color = texture(texture0, fragTexCoord);
    
    // Blur calculations
    for( float d=0.0; d<Pi; d+=Pi/Directions)
    {
	    for(float i=1.0/Quality; i<=1.0; i+=1.0/Quality)
        {
	        Color += texture( texture0, fragTexCoord+vec2(cos(d),sin(d))*Radius*i);		
        }
    }
    
    // Output to screen
    Color /= Quality * Directions - 15.0;
    finalColor =  Color * multiplier;
}