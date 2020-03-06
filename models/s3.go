package models

type S3Policy struct {
	Version   string
	Statement []S3PolicyStatement
}

type S3PolicyStatement struct {
	Sid       string
	Action    interface{}
	Effect    string
	Principal interface{} // 用户
	Resource  []string
}
