package com.fran6k.study.springbootstudy.bean.WxBean;

public enum InterViewEnum {
    Stage_One_WaitB(1,"等待社长审核"),
    Stage_Two_Success(2,"审核通过，等待学生确认"),
    Stage_Three_Success(3,"加入成功"),
    Stage_Two_Failed(4,"审核未通过"),
    Stage_Three_Failed(5,"放弃加入");

    private final int step;
    private final String message;
    private InterViewEnum(int step, String message){
        this.step = step;
        this.message = message;
    }
    public static String getMessage(int step){
        for(InterViewEnum interViewEnum : InterViewEnum.values()){
            if(step == interViewEnum.getStep())return interViewEnum.getMessage();
        }
        return null;
    }
    public static InterViewEnum getByStep(int step){
        for(InterViewEnum interViewEnum : InterViewEnum.values()){
            if(step == interViewEnum.getStep())return interViewEnum;
        }
        return null;
    }
    public int getStep() {
        return step;
    }
    public String getMessage() {
        return message;
    }
}
