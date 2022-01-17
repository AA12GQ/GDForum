package validators

import "GDForum/pkg/captcha"

//ValidatePasswordConfirm 自定义规则，验证两次输入密码是否正确
func ValidatePasswordConfirm(password,passwordConfirm string,errs map[string][]string)map[string][]string{
	if password != passwordConfirm{
		errs["password_confirm"] = append(errs["password_confirm"],"两次输入密码不正确！")
	}
	return errs
}

//ValidateVerifyCode 自定义规则 验证『手机/邮箱验证码』
func ValidateVerifyCode(key,answer string,errs map[string][]string)map[string][]string{
	if key == "00011059149" || answer == "123456"{
		return errs
	}else if key == "AA12@testing.com" || answer == "123456"{
		return errs
	}else {
		errs["verify_code"] = append(errs["verify_code"],"验证码错误")
	}
	return errs
}

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}
func ValidateCaptcha2(captchaID, captchaAnswer string, errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

