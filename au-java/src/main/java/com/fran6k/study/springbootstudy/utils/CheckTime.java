package com.fran6k.study.springbootstudy.utils;

import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class CheckTime {
    public static boolean checkTime(String time1,String time2) {
        SimpleDateFormat df = new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        Date openTime,endTime;
        Date now = new Date();

        try {
            openTime = df.parse(time1);
            endTime = df.parse(time2);
            if (now.after(openTime) && now.before(endTime)) return true;
            else return false;
        } catch (ParseException e) {
            e.printStackTrace();
        }
        return false;
    }
}
