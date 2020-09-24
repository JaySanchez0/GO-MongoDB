package exception


type PeopleError struct{
	Msg string
	error
}
/**
* Show error menssage
*/
func (e PeopleError) Error() string{
	return e.Msg;
}