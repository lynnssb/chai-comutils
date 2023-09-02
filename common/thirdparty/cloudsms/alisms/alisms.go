/**
 * @author:       wangxuebing
 * @fileName:     alisms.go
 * @date:         2023/5/19 18:17
 * @description:
 */

package alisms

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @return error
 */
func createAliSmsClient(accessKeyId, accessKeySecret string) (result *dysmsapi.Client, err error) {
	config := openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	result, err = dysmsapi.NewClient(&config)

	return result, err
}

/**
 * 发送短信
 * @param accessKeyId:
 * @param accessKeySecret:
 * @param signName:        短信签名名称
 * @param tmpCode:         短信模板CODE
 * @param tmpParam:        短信模板变量对应的实际值。支持传入多个参数
 * @param phoneNumbers:    接收短信的手机号码(支持对多个手机号码发送短信，手机号码之间以半角逗号（,）分隔。上限为1000个手机号码)
 * @return SendSmsResponse
 * @return error
 */
func SendAliSms(accessKeyId, accessKeySecret string, signName, tmpCode, tmpParam, phoneNumbers, outId *string) (*dysmsapi.SendSmsResponse, error) {
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	request := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  phoneNumbers,
		SignName:      signName,
		TemplateCode:  tmpCode,
		TemplateParam: tmpParam,
		OutId:         outId,
	}
	runtime := util.RuntimeOptions{}
	response, err := client.SendSmsWithOptions(request, &runtime)
	return response, err
}

/**
 * 批量发送短信
 * @param accessKeyId:
 * @param accessKeySecret:
 * @param tmpCode:         短信模板CODE
 * @param phoneNumberJson: 接收短信的手机号码(JSON格式)
 * @param signNameJson:    短信签名名称(JSON格式)
 * @param tmpParamJson:    短信模板变量对应的实际值(JSON格式)
 * @return SendBatchSmsResponse
 * @return error
 */
func SendBatchSms(accessKeyId, accessKeySecret string, tmpCode, phoneNumberJson, signNameJson, tmpParamJson *string) (*dysmsapi.SendBatchSmsResponse, error) {
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	request := &dysmsapi.SendBatchSmsRequest{
		PhoneNumberJson:   phoneNumberJson,
		SignNameJson:      signNameJson,
		TemplateCode:      tmpCode,
		TemplateParamJson: tmpParamJson,
	}
	runtime := util.RuntimeOptions{}
	response, err := client.SendBatchSmsWithOptions(request, &runtime)
	return response, err
}

/**
 * 查询短信发送统计信息
 * @param accessKeyId:
 * @param accessKeySecret:
 * @param isGlobe:         短信发送范围(1:国内短信发送记录;2:国际/港澳台短信发送记录)
 * @param startDate:       查询的起始日期，格式为yyyyMMdd
 * @param endDate:         查询的截止日期，格式为yyyyMMdd
 * @param page:            当前页码。默认取值为1
 * @param pageSize:        每页显示的条数。取值范围：1~50
 * @param tmpType:         模板类型(0:验证码;1:通知短信;2:推广短信(仅支持企业客户);3:国际/港澳台消息(仅支持企业客户);7:数字短信)
 * @param signName:        短信签名
 * @return QuerySendStatisticsResponse
 * @return error
 */
func QueryAliSmsSendStatistics(accessKeyId, accessKeySecret string, isGlobe *int32, startDate, endDate *string, page, pageSize *int32, tmpType *int32, signName *string) (*dysmsapi.QuerySendStatisticsResponse, error) {
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	request := dysmsapi.QuerySendStatisticsRequest{
		StartDate:    startDate,
		EndDate:      endDate,
		IsGlobe:      isGlobe,
		PageIndex:    page,
		PageSize:     pageSize,
		SignName:     signName,
		TemplateType: tmpType,
	}
	runtime := util.RuntimeOptions{}
	response, err := client.QuerySendStatisticsWithOptions(&request, &runtime)
	return response, err
}

/**
 * 查询短信发送详情
 * @param accessKeyId:
 * @param accessKeySecret:
 * @param phoneNumber:     接收短信的手机号码
 * @param bizId:           发送流水号
 * @param sendDate:        短信发送日期，支持查询最近30天的记录。格式为yyyyMMdd
 * @param pageSize:        分页查询，每页显示的短信记录数量。取值范围为1~50
 * @param currentPage:     分页查询，指定发送记录的的当前页码
 * @return QuerySendDetailsResponse
 * @return error
 */
func QueryAliSmsSendDetails(accessKeyId, accessKeySecret string, phoneNumber, bizId, sendDate *string, pageSize, currentPage *int64) (*dysmsapi.QuerySendDetailsResponse, error) {
	client, err := createAliSmsClient(accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}

	request := dysmsapi.QuerySendDetailsRequest{
		BizId:       bizId,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		PhoneNumber: phoneNumber,
		SendDate:    sendDate,
	}

	runtime := util.RuntimeOptions{}
	response, err := client.QuerySendDetailsWithOptions(&request, &runtime)

	return response, err
}
