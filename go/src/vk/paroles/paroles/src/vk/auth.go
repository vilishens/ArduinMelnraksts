package web

type pass struct {
	name string
	pass string
	salt string
}

func init() {

}

func Secret(user, realm string) string {
	if user == "john" {
		// password is "hello"
		//return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"

		return "$2b$10$i11qD/kMPQzGeHtiG7Jad.Gn1rIxs.MK0ohp9agSwp.wx1wIPgz1O"
	}
	return ""
}
