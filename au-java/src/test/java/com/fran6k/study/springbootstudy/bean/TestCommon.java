package com.fran6k.study.springbootstudy.bean;

import org.springframework.web.client.RestTemplate;
import org.testng.annotations.Test;

@Test
public class TestCommon {
    int a = 0;
    public static void main(String[] args) {
//
        RestTemplate restTemplate = new RestTemplate();
        String url = "https://mp.weixin.qq.com/s/m_90_3fNTr6Ingz2nk2Y2g";
        String body = restTemplate.getForEntity(url,String.class).getBody();
        assert body != null;
//        解析时间
//        int index = body.indexOf("if(window.__second_open__)return;");
//        int last = body.indexOf("e(t,n,i,document.getElementById(\"publish_time\"));");
//        String test = body.substring(index,last);
//        test = test.split("\"")[5];
//        System.out.println(test);
//        解析标题
//        int title_first = body.lastIndexOf("<meta property=\"twitter:title\" content=\"");
//        int title_last = body.indexOf("<meta property=\"twitter:creator\"");
//        String title = body.substring(title_first,title_last).split("\"")[3];
//        System.out.println(title);
//        解析图片
//        int begin = body.indexOf("var msg_cdn_url");
//        int end = body.indexOf("var cdn_url_1_1");
//        String picUrl = body.substring(begin,end);
//        picUrl = picUrl.split("\"")[1];
//        begin = picUrl.indexOf("\"");
//        end = picUrl.lastIndexOf("\"");
//        picUrl = picUrl.substring(begin+1,end);
    }
}
