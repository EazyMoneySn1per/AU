package com.fran6k.study.springbootstudy.config;

import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

public class InterceptorConfig implements WebMvcConfigurer {

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(new JwtInterceptor())
                 // 拦截的请求 /xxx/**   表示拦截xxx下所有
                .addPathPatterns("/account/**")
                .addPathPatterns("/activity/**")
                .addPathPatterns("/admin/**")
                .addPathPatterns("/AMinister/**")
                .addPathPatterns("/dataShow/**")
                .addPathPatterns("interview/**")
                .addPathPatterns("/outlay/**")
                .addPathPatterns("/synthesize/**")
                .addPathPatterns("/twitter/**")
                .excludePathPatterns("/mpapi/addWxUser")
                .excludePathPatterns("/auth/login")
                .addPathPatterns("mpapi/mytest");
                 // 不拦截的请求 :用户认证、用户登录
    }

}
