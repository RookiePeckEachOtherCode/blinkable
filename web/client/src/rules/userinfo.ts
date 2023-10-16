
import type {FormRules} from "element-plus";
export const rules:FormRules=({
    username: [
        {required: true, message: '请输入用户名', trigger: 'blur'},//是否需要校验，鼠标失焦时触发校验
        {min: 3, max: 10, message: 'Length should be 3 to 10', trigger: 'blur'},
    ],
    password:[{required: true, message: '请输入密码', trigger: 'blur'}],
});