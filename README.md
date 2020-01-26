# Openmock Customization Example

This is an example of how to extend [openmock](https://github.com/checkr/openmock) 
with customizations, such as a custom handler for kafka messages.  To do this, 
you copy the generated swagger 'main' from openmock into your project, and modify it 
to create your custom `OpenMock` instance, then pass that instance in to the `ConfigureAPI`
step. 