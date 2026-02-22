package module_say_hello

// Untuk menghindari adanya fitu yang tidak kompatible pada versi module terbaru/terlama
// lebih baik mengganti nama modulenya dan dengan tambahan versionnya 
func SayHello(name string) string {
	return "Hello " + name
}