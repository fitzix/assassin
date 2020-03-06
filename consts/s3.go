package consts

// S3PolicyAllowImageStatic (ARN) Amazon 资源名称 唯一标识 AWS 资源
// arn:partition:service:region:namespace:relative-id
// 协议:分区:服务:区域:account:bucket名称/资源路径
// refer https://docs.aws.amazon.com/zh_cn/IAM/latest/UserGuide/reference_policies_elements.html
const S3PolicyAllowImageStatic = `
		{
			"Version":"2012-10-17",
			"Statement":[
				{
					"Sid":"AllowImageStatic",
					"Action":[
						"s3:GetObject"
					],
					"Effect":"Allow",
					"Principal":{
						"AWS":[
							"*"
						]
					},
					"Resource":[
						"arn:aws:s3:::%s/images/*"
					]
				}
			]
		}
	`
