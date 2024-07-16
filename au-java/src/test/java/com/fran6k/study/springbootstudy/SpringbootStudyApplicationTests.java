package com.fran6k.study.springbootstudy;

import com.fran6k.study.springbootstudy.bean.WxBean.InterViewEnum;
import org.springframework.boot.test.context.SpringBootTest;

@SpringBootTest
class SpringbootStudyApplicationTests {
    public static void main(String[] args) {
        InterViewEnum[] interViewEnums = InterViewEnum.values();

        for(InterViewEnum interViewEnum : interViewEnums){
//            System.out.println(interViewEnum.getMessage());
//            System.out.println(interViewEnum.getStep());
//            System.out.println(interViewEnum.name());
            System.out.println(interViewEnum.ordinal());
        }

//        System.out.println(InterViewEnum.getMessage(4));
//        System.out.println(InterViewEnum.WaitOne.name());


    }

}
