package com.fran6k.study.springbootstudy.config;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fran6k.study.springbootstudy.utils.JwtUtil;
import org.springframework.web.servlet.HandlerInterceptor;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.security.SignatureException;
import java.util.HashMap;
import java.util.Map;

public class JwtInterceptor implements HandlerInterceptor {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if(request.getMethod().toUpperCase().equals("OPTIONS")){
            return true; // 通过所有OPTION请求
        } else {
            //直接获取，前端token不需要加上Bearer
            String token = request.getHeader("token"); // 获取请求头中的token
            Map<String, Object> map = new HashMap<>();
            try {
                boolean verify = JwtUtil.verify(token);
                if (verify) {
                    return true; // 通过验证
                } else {
                    return false; // 未通过验证
                }
            } catch (Exception e) {
                e.printStackTrace();
                map.put("msg", "token无效");
            }
            map.put("state", false);
            // 将map转为json
            String json = new ObjectMapper().writeValueAsString(map);
            response.setContentType("application/json;charset=UTF-8");
            response.getWriter().println(json);
            return false;
        }
    }
}
