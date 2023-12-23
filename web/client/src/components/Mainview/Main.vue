  <script  lang="ts">
  import $ from 'jquery';
  import {Lock ,Message,Notebook,Comment} from "@element-plus/icons-vue";
  import {Getcards} from "@/apis/getcards";
  import {Likeapi} from "@/apis/likeaction";
  import {guestbookApi} from "@/apis/guestaction";
  import {useUserInfoStore} from "@/stores/userinfo";
  import router from "@/router";
  import {ElMessage} from "element-plus";
  import { ref, defineProps } from 'vue';
  export interface Admin {
    admin_id: number;
    image_url: string;
    signature: string;
    icon_url: string;
    title: string;
    gestbooks:Guestbook[];
  }
  export interface Guestbook {
    admin_id: number;
    book_id: number;
    context: string;
    avatar_url:string
    create_time: string;
    user_id: number;
    user_name:string;
  }

  export  interface Guestbookfrom{
    user_id:number;
    admin_id:number;
    context:string;
  }

  export default {
    setup(props, context) {
      const guestbookfrom = ref({
        user_id: Number(useUserInfoStore().getUserId()),
        admin_id: 1,
        context: 'cin',
      });
      const whoguest="admin"
      return {
        whoguest,
        guestbookfrom,
      };
    },
    computed: {
      Notebook() {
        return Notebook;
      },
      Comment() {
        return Comment;
      },
    },
    methods: {
      hideButtons(id: number) {
        var type=false;
        $('.hide').each(function () {
          if (this.id.startsWith(id + '-'+'1')){
            if(this.style.display!='none')type=true;
          }
        });
        $('.hide').each(function () {
          if (this.id.startsWith(id + '-')) {

            if (type) {
              $(this).fadeOut('slow');
            } else $(this).fadeIn('slow');
          }
        })
      },
      hitarrow(){
        // 获取目标元素
        var targetElement = document.querySelector('.card');
        // 检查目标元素是否存在
        if (targetElement) {
          // 获取目标元素的位置信息
          var targetPosition = targetElement.getBoundingClientRect().top;
          // 调整目标位置（减去 100px)
          var adjustedPosition = targetPosition - 150;
          // 使用 scrollIntoView 方法滚动到目标元素
          window.scrollTo({
            top: adjustedPosition,
            behavior: 'smooth',
            block: 'start',
            inline: 'nearest'
          });
        }
      },
      tableview(id:number){
        if (id == 1) {
          console.log(this.table1);
          this.table1 = !this.table1;
        } else if (id == 2) {
          this.table2 = !this.table2;
        } else if (id == 3) {
         this.table3 = !this.table3;
        }
      },
       async likebut(adimin_id:number){
        const user_id=Number(useUserInfoStore().getUserId());
        if(!user_id){
          router.push("/login");
        return;
        }
        const form={
          admin_id:adimin_id,
          user_id:user_id,
        }
        const res=await Likeapi(form)
       if(res.status_code===0)  ElMessage.success("点赞成功")
         return;
      },
      async guestbookbut(adimin_id:number){
        const user_id=Number(useUserInfoStore().getUserId());
        if(!user_id){
          router.push("/login");
          return;
        }
        const form={
          admin_id:adimin_id,
          user_id:user_id,
        }
        const res=await guestbookApi(this.guestbookfrom);
        if(res.status_code==0)ElMessage.success("留言成功")
        return;
      },
      async getcard() {
        // const res=await Getcards();
         let p:Admin[];
        // if(res){
        //   p=res.data.admins;
        // }
         p= [
          {
            "admin_id": 1,
            "title": "crk算法第一臭卷直男",
            "signature": "哼哼哼啊啊啊啊啊啊啊啊啊啊啊",
            "image_url": "data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAsJCQcJCQcJCQkJCwkJCQkJCQsJCwsMCwsLDA0QDBEODQ4MEhkSJRodJR0ZHxwpKRYlNzU2GioyPi0pMBk7IRP/2wBDAQcICAsJCxULCxUsHRkdLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCwsLCz/wAARCAEKAcgDASIAAhEBAxEB/8QAGwAAAgMBAQEAAAAAAAAAAAAAAgMAAQQFBgf/xABJEAACAQIDBQUFBAcFBwMFAAABAgMAEQQSIQUTMUFRImFxgZEUMkKhsQYjUsEzYnKCktHwFUOisuEkNFNzs8LxFnSDJWN1k9L/xAAUAQEAAAAAAAAAAAAAAAAAAAAA/8QAFBEBAAAAAAAAAAAAAAAAAAAAAP/aAAwDAQACEQMRAD8ARu1PM0YhHU1s3uFF7xt6USzYIcYmoMYgX8RqbgDnW/e4E8I2FUwwre6SPKg54i8asRmt6xwjW5ot3EaDGsbGnJh3PO1aFjjHEm1NCwDkxoM4w9udQwnrWrdxNqMw86IQLYdo0GHdOam6Yc63boAaE0JjNBlVD1pojNuNPWG/GniJbd/hQYxG1GI25GnlLVagjpQJET1RSW9grHyrWMnNhTAydRQYck3NDVFG/Ca6QZTwtQMl9bCg5hVu/wBKsK3fW3dG/u1YjkB0T5UGTLJ+HSlm/wCE11As/Ax0t425oBQc/X8LVdieRrZuWPIUO7I6UGXI/Q1WSW+gPpWrdvrahO+XlQZwkvQ+lXZ+YPpTwZuNMDNzUUGYB+h9Kna6H0rWMQq37AvSzipG+FAPCgRqeRqijH4W9KcZpDzUeAqxicQOBW3hQZTGeh9KExt3+la/bJxyT0oGxkxHup6Cgy7s99XZu+jbESfhHpQGaXlp5UF/ef0KlnP/AIoQ8x5mr+9PWgvK/Sr1HEVAsvU0Yjc8WoFFu6gLmtIw9/iFQwAcxQZszVMzVo3adRV2j7qDLdqo5612TuqiF6CgyZX/AKFTLJ/QrUSByocx5CgSFk6H0q8j9DTc7DlV71h8NArI3fUyNTd8TyHpUzk8qBWQ1eS/OmanlV5RzNAsRnrUptkHxCpQGYI9SF9aErEo1C/KjfD3JvP86T7HETczXPcTQWEiY8Fv5UZjQjgKD2aIcJfrTVgjt+lNAGQdBRpHGeLCjESjnfxoiijW4oCEUQtaxo9zGRralBwPhq9454LagYMMvImmezjkDSRNiF5U5cS494GgsYaiEKj4SantYAPZofa3PAWFAYjH4DRAINN016FJ2Jua0rPYcRQKzJziqmMfOKwq2dmY2pLyS6rQUfZb3N/SiEuBA90+lIzOKBnJv2R6UGwzYHlcUlpoATYm1ZSWPwiq8hQafaIhwY+l6r2luR+VIFx8K0YZvwLQE2JmPxN5ClGYn3manDM3wrQNCx5CgQZBf3mqt73mmmE916AxAcbUCzI3U0JdvxGm5F7qEoO6gXnb8RqZ26mjyjuq8o6CgWWPfVXPfThbpV9noKBFz31RduppxA6UBHdQKzHrVZj1o8t6m7FAvMalzTMlTJ3UABjRB2HM1eXuogB0oKzt1q87Vdl6Vdl6UA5261RdjzorDpUsOlAvM1VdqZYVVhQDmapmartVUEzGpnPSppUoLEh6Cr3p6Cl6VWlAZlH4RUEvcKWbVWlA7e0OcGlVdqAiwqUNjUoNZmj5rx770O9hHKsJc60JJ60G/wBog6CmCeE87eArlg0YNB1BNH+KjDxn4h51zA1MD0HUUoea0YB5FflXLDnrRiWg6lm5j6VMluINc8SsPiNGJ307RoNtkHFT6VMqW0Vu7SshxEnDMasTy2sGsKDQQbc6VmYczSy7Hix9aIG/OgLfTC9qm/k/BeoCelFny8VNABnPAp8qAyH8BpjSoeC60G8J0oB3v6pqt7+rRanlUy8jzoBEw/DRe0AfDRbod1CY/D0oK9oPIVDiZOVqEx91TdHpQEMQ9wdL1bYgva9tOgoNyam5IoIWXrQFhRbq1UY++gHMKmfvqFKorQXm76q561VqlvGgmvWprVW8alqC9KsFelDarymgK4qaUOVquxoL0qWWqsauxoJYVVXY1LHpQVVGitUoA1qqI1VBV6q9QiqtQS9VertVWoK1qtaK1VagA1VGRVWoB1q9alqlBNalSpQZi2tVeqYWJ1BoaBgogaUDRg0DQaIGlijFAYNEDSxRigYDRA0sUQoGg0QNKBogaBhNQMRQ3qUDM9Qsx50uiBoJc1AauwqrUBhj0qiSeFCKIGgIOeBNTMOpqzlNAVoDuDwNVmI50ABqzcigLO3WpnNBrzqfOgPNVEg0NUb0BWBq8lBejVrUE3dTIOZotDUoKyLVGNau9XQBuxVZKbY1Mh7qBWWry03dtV7tulAnLV5TTt23SplPSgTlqWppU9KrK3SgVY9Kq1Osaq1AkrQ5aflFCVoE2FVlFOy1WWgVaqy03LVZT0oFWqiKdl7qHL3UCSKq1OK91Vl7qBVqq1Ny1MtAq1Sm5alByuZqwKmt6IUFgUYWoKMWoKC0YFXVigoA0Qq6sUEFXVgVYAoKFEKlhVgUF1dS1SgupV1dBL1L1KlBKlFYVLCgq9XepYVLUEvVVdqlBKGjtUy0A1KK1XagC1SjyiplFAIvV60WWitQLqxR5RVhRQQLeiK2qwBRC3SgWLg86K/jRECqsKCa8qq7UY0qiBQUNalh1qWqWoKKnlVZTRVLmgDLVFe+jN6rWgDL4VMpo7VKBeU9KG1PqiAeVAm1SxpuUDlVWFAkg1VqflFTL3UCNalOyjpUyLQI06VKdkFSg8/zoxUI1NEBQWKMVAKICggogKgFGBQUKKoBV2oLFXUFEBQVVirtV2oKoqlqu1BLVdqgFFagG1XaiAq7UAWqUdqlqAKu1FapagGrtRWogtAFqu1MC1MtAu1Xajy1eWgDLRBaILRBaBeWry0dqlqALVdqK1WFoBqWpiozEBQSTyFMOHkS2ZCPGgz2q7U4xkcqrLQLtTocM8p0oo4izAWrs4aFUUaa0GAbMJANzfuGlZ5sHLHckGw7q9IFPIVmxkTGNr31Fhbr30HmytUVrU0ZBsRbxoCtBmy1MvdT8lQJ3UCMtTLTyvdVxwyytljUse6gz5aq1dL+zMba5QDhzB+lZXiZGZWBBBtrQZ8tVlNPy1MtAixqWp2WpkoE2vVWp+SqyUCctSn5KlB5zIb0QWs5wUtzcn0NEMJJwDH5/wA6DSFNEFrOMLJ+I+OtX7NKPiY+tBpCmjC1mWCUDW/+KmR4XEzOFjVi9r5QTyoG5TV2rXJsHaYSFoSk28Un7pmGUqLlWvWY7M2ioBaKUXFwSr6+FBAKILSWwWMjyllcBvdLXs3hVrh5uZa4460DwKsLSRBNe350z2eXTX5mgPLUC0vcyg8T6mr3Mp5n1NA0CrC0tcLiNSc4XmbGwA61yJNrjdN7PG++bOqmQ6R24OOt6DvBDRZK8oMdtdwAcXKvPsBU/wAgrfh9pzlyuJYlXIyuo9zlY25UHcy1MpoRFKVDAsQQCCLkEHUEGoIpuea3nQFkNXu6sRv1PzqzG4HxfOgHJVhKHdyfrfOmJG3PN86CwtXlpow7EAjN60DQuOZoBy1MtFupOjfOnJhMVIFKxvlPOx4daBGWplp2IwkuHJMjgIFDlybADqTQywNEmGd5EyYlssJVrhmtm5UC8tTLVosbmwxEOugvKgH1rcuzHkXNFNHJa4YI6tYg2Pumgw5aILWt9m4qNQzcD33rmy4pIHeIXeVGykH3LjjrQdHCoA4a3A12wYpAAyjMBa9hwrycG1JUYZ4VZAQTuyVa3dc2rqQ7Yw0hZWWSK/us5BHnl4UHSxGGjlAAUAqLAgAX9KwHC2Go1pox+FLBVmQnlrxPdVT4plWRgFuqM/DiVGaguKDLY21rYoyc68zDtHHCZZJJnZSw3i6ZSvOy1tG2AZHDIBFrkIvn04X8aDvq558KaGVhlYaGvFz4vEzy70ysuVgY1VjaO3Stb7ZxQjh3TBZQ53hYBhYDgL8jQejxODinUZQFdb2PXuNcwYM5irqbi40PPlWfCfaAxoyY1XkIIyyQql7H8QJArWm1dlM2YyyKWsSZIjYd5K3oFtgyuZeOmlvWphoFRiX5cL9a1YjF4KOGTEGdGjQqv3dnLE8ALVjGKhnQPE1wTfoQe8UDmhw0rXK5TfiugPlWqCCKMWQi5NyedhXHOPw4mfD5yJFuGJHYBAuRmpvtT3BVuAGUjpQdwSG9iLEcOlZcfhlkgLKi51INwNe+ssePkNgw8xW1ZlkiKk2PO/GxoOAVtp0qWrsnCYWYACPISbh00PnXNnwU0EhU9ofCw4EUGey91QlBzosqn4k4X1IqZUJADoSdQAwN6ACyd/pVZhyBpojPCmR4XESHsxnuJso9TQZsw/CalbHwGLRc7Rac7WJHjapQZZlwS8RHckKMq6+Fs9WFjcKrLh+woVN4scbg8bFibmsMuMxBsGkZR03IHzAovbM+QM6sVXKL3JHfrQdM4eIR39jUnshjHaax62T+dLePBoqtuRlC3kuhyp32tmFZlnkPuommmYC1WcYysbuyvzyuQP4TpQM/2ZxeOPDuWPZEcqAeh/nRHexqCIcik5d4iOwv0Yre1JXHgXOWMkm+Z1jbXzq/b3vcmNRe43aooAta2hoHhNpgZEEgHEgG3GkbQxv2rhTDx4HBupyM88jRxzODewRUFwOvCo+NzsTvFU8BrqRawuQRTotobSKgJig5txZYSflY0HLXaH2rnaJZ8NghGjLdZ8PJDa/xEgg10IcPtNw8ksuDZ5S5GWBdBe9lBH1o5J9q4i2aYEXvaOGMa9zXvTM+0wVymbRieyqHtdCL0ADFyRb1Xw2FW1rGFdWI6liRSX21GoBmwA1zZTC6DNbTtWtanyGZz99G3LjBlNZmiia4LwopBuDGy/Mmg5L7Y2sQ+SPCITm1EJYi/C2duXhSBj9q+zmDfNnMpk34sJypH6MsOXOuucHhW0EsPEasrDx1vVrsvBHU4mMa6gN9OdBwhiNrA3ONxbdlkIkmkZSrDKQQxtWZcIOFq9Q2y4F4Srx+K1reINLGAw6+/LFbojpf5mg4K4YAC406nhpQmEXuBXohDs3gFJ5ak6fOjGE2fe+4Y6aKS5166Gg4azY5RGq4mcLELRIrsEUdABpWr+0NpG5Mqm4Ye4gtcWuMo49KvbW0dg7BwonxOGaTESA+y4dGdWlKntG7k2QfFXmdm/bPAy4iUbR2akcDspjGFeQmK+nAnXv1oPUrtLaIyapoANYxr3nvq5cZtGaO2cImYX3Iy+pGtdKFtnukMsECSI6gxOoEgKnoTenmWCxQwgKeIyqL+QtQcnDYqSObDGdmeEALIptcgc7nnXckxWx0V2QwPkjLsiy9sn8I141gZNlk9qCMHvdx/wB1WBs6wtBhjYcTqfmaBUm3Jikaw4SCIq+ZrkyB1sQEIb50Y2yXhkvhlXEEEJJHZkF9LlXpu8wNrCDDW8FI+tBm2fygw4P/ACxQcsyYj/jTG/WRv50ftW0OypxE5UcAZG4HpXTz4a9xHBz/ALtfrV3wn/Bw58I1FBypZsZIhiaeYxGx3bSMU01GhNLlknlWBHPZgTJENdB18a7V8CeOGhv3ACivgD7+GjPiATQec7Sk2qw0wIKkg8AVuD6ivQf/AEy99xHfoQKu2zQMvs0Nu+1BysNtDauHzbvFSgNxV2zj/HSgNeHfXaDYFdRh4RbmMtX7RhuBRbDgOzb50HISKWQ2RGY9FBJ+VQBgeHDyrspi8PGcyLkbqlhfxtQtiMC57cak8dVB+lByuWnzomlnZSDIxXgRmJFdP2rAr7qEW/CALfKhM+zyQSjH91fzFBytKlr866u/2bzia/7CfyqziMKQMqMLajSMEedqDkcOdWq3te9jpeuoJ00tbThcp/8AzQXw2p3KlibmxX8hQc7KxvlUm3GwJ+lXHBi5iVhhlc2uRGjNp5V0xiHQgpdeXZJ/KqOJmN7u/P4iOPhQc10xUSmJ1kQEgsjqRdl0vY0UM08GbIoIa1wQdSPCtplcgA9oDgH7QHgDQDdBg27Fwbg66HqKDmFWzMTe5JJ8+NaY550UKp0BIW9j5a1sy4VtWi535nX1pgGD17Ca9U19aDEmNxcd7FSb/Eo0py7RxhFyBYnuGY91OMOBb4U8riq3ODvcC3gTb0oLi2xtGNmC7mxvlEik5e++alSbR2w62bEShSuXshVBAuL3UetO3eF0Fgbdf9asx4crl1C9ATQcwAiiFxroPrXQEWGFuwvjbX1qmiw7ADKbAm1jbjQZ45GQow99TcNc30rb/aWLY27FuAuGNvnS93Be5XXTXhwFuVWVg/CPGg1RbXxEXYmVW11tobd1Ss5dOg07galB5HLKCbPceB/nU+8B1k16W0qGfDX0nQ+Gv0olnhOgkF/Aj60FgyAauCOmU/nV2c2Oc28LUxZIgCWlVdL663vy0o1YNmMckZCrmJugt36m/wAqBYAPEk+Bowi3vZgfEj6VRljW95UOvIqfpUM8Q/vBfyoGBQPjcedx86YFc8HJ8v5Vm9uw0Z7cmU8NRf6VpTE4JkzDExqCCeDE+eThQWBOPjPce1pRWxB4SuD1zN+dAMZhRxmS1r3va/rTfaYMmfeLkva5NqAcmIP98fM0Qja1mb0NqntmDsTvAwUXOUHT1pa7S2eSBv110JNk/wA5FA3dDq3qaIR26n+u+q9swANt5+ySRZvCmifDtorENx7Yyi3idPnQL3a8LGpuV76cWCgkgEdVdCD6Xq1YMuZY3YG/uFH4cb5T+VArJ+1VGM/iejGKw3NiD0Km/wAqYjwyAFZFF/xELYdSCaD519to5jtPBo2dlfAq8ehJNmfMot0tXn4ITIyxQqzztkSJB8TA3JPdX1XaezsJLNg9qK0cj7J9olnkjljKDBmCTeKTc37gNa+axTSzSythbYZbuYIU7eRCcwhzntG3DWg+kbLMSwpswYlTjcMjzSRLfNu2feXycgL2sfHXNW/I446nrXg9l4ub2OSXDyzri8I0Rebexxow3jyxxyW7ZGrrfW3Dn2fW7I2su0ZJsNMzDFRqkouixkxyAkAhSRcag+HfQb8t/hHmKrda3t6CtBRRe7gW43a1LvGb5XVranK5ItQAI9NVv42q923IW8DTAY7aMhPTeAmiXITp/p60Cd23xfWryEdfWtBQc1T1BvVGM2vkUBbai2npQJyk8QarIR8Pzo8yE2DrfprRlba5ltw4g0GfK/4amWTp9adYnQMPUX+dHupQL2NvEUGezaaGiGb+hTckv4vmKAlhxfjy40A2PIfKryt09CasFWNs8d+eoFGYiti6kDrmAv60C7H8J+tXb9UelGI2cdkkjxB+lQRycj86Aco/Aauy9Got3Jzt5kVN3JytfucfzoBOX8NVcDgp8jTcrEasNOrKbVWQHmtvL8qBWY9G9avMelO3I/oiqMBGvLrQJue6q7XSmhBr2lPmL+lTKBzFAu7Vd2owoa9iunU2+tGImJsCnd2hQKue75VV37qfuH6p/EKsQH8aDzP8qBAzVd2p25W+rqD5VNweTJ/GtAq5qrn+jWn2bs3zrfoNfpQnD8PvE15E2PzoEXPT51Vz0+dahhHIzBkt3G/0pZgbx8P9aBGZulSmmMA2OYHvBFSg8c2y8SpIVomtxtPEe/UK1D/Z8qL21ytcf3sVgOhCm/zr1JwDSF2jyuSBdnijXh3hB9aBsFIgPag5BuHZ7sutB51dnTPpCxYgaiKSxseuU0ceyMYr5jEZlbskSxsQL87+8LeNegTCYYks7Rx5VB+7Ew860w4PZosZHixOnY32Ylbcxfs0Hm22Wkci75wsbuQFfPGCOiSMaXJsrDoWMWJkHAWAEoPddLfSvVSYPCyLIEMIj9903quh5A5WNqwnA4ZWK2ZAvvf7Oy5e8rc0HAXZuHYFWmuxsFLLLHbw0t86Ntj7RgIKrMBbR4QzKQe9K7ZwqZyI5C4Pu5Vdb+CnSuhCZYYlhdSOzYNny3F/hIPGg8iMDjGY5pr9c2Yse43Wt2H2XhplRcUXiI0EkCOb95AP/bXpGjMke79ocLxtLKjEXHNnANY/ZokYAyKRpbNMtmHdY0GBvs/BkzYbFyzHS6mPXj+varXArDo63slrSQKjm/6yhvyrsgubpvsNGjclne4B7gbUAwsLtZZ45JB8Kl3BPjcUHBGzo2Ybp8MpY3JYMhXwutdSBfZlCvhZJCLZpIirq9uZGtbWRADmOGABudTdMvG1iaZF7OVOVw/WzWt/EL0GJXwpGSbZ2Eyk3zLvYJhz1IFq1xHCpbcYdVW4yk2OX9bQXo2OEkyh3gNuCmSJn08GvQGLCqM6GNkJy6Jmy8uZvQeGk+1213favtuzREsE+ISKEJIohjiBsMQyF5OmphAPUV39k4/Y20MNCDicBjZZUiaTDxkrunKBzGAGDm3ffhXXxWA2diTGcVBhZFTtAzw5ittARnB0rwv202VsrBDBYnCGBExhfDYqDCfcvI8TidZrRjJpoL26UHX+2uMwuD2I0GDwyRSY7FRYSR7BZFhRTK6qTrrYA+P61fNGDBBJHcCNbycMykngSK0TYrEYtcNHPjcbicJglleFMZJnMOcqCFJ62HLlS1ciGfLa80cjSjQqEtoV+hoOlhZpcHBtLIyouJwUJC3KXR5Yzfs69fWn4TGY3CS7JxblWZcZisI7kSF5lyJLGjyc7X06eFLSOPEpsGEGPNi9jDDa3JM4ZwlwuupAW961PG+L2HsmfIIXwmNXDb665icOsjE5SePbj9KD6FhpMLjsJDPJJczxq+VkDZb8gRThgEA7IjC8C2YroeVntXnPsthttwQTPtCCSbDYmTe4FUkjjQoSTvxm5NpoDyr02RBZmw80Q4jVJG88jUGd8FFZcii55rJf5Gr9imUEhb87gt6GukkkAQgo1jydXBsOmYW+dZ5G2YDqwQnjmBH0FBjTDtm94ob/AB3I+VMeOICzySEaAELlt52FOaTBR5SQzIbe8CLeRsave7OPu4gR5tQulh3do0GPdT3+7mDDkpNzbvzCoI5iBdEB4392/neurGuFlF94j3sBm3an5NVvhsMlyGjQi3ZeRVv4ZqDjvDILby2tzo2Y277a/OgWGO5KSFGHKzMD867BWEhcr4c349rL86E4ZSQLBs3Dcsh+dBydxITYS4drG9nUjXx4/OrWOZCT2bfqkgfOun7KuaQNDKG5BmQ38yRVCGBL7yGUC9uwSRfyJoOWYiSLRk36MLVZjYWGQgcO1c/MGutuMKTYbwaXNnA18Gsahw2GvbPoPxk3PhYUHGMaAgAuTbkdPnrTBG3IOB36+eldRo9nRjt7xjyKpmX63+VS+AAACsqniWJA8wDQYVDAauLjgG1B76o7zX3CDrpl+hFbj/ZZuFmUka9lXPzIq09gPF4we9jfzBAoMKxg8QvDmB+VGY1W4QIb82zgjwsa6YjwhsRLED+slx6qxqSRwKoZVgfrYn5XIoOaqYge6LD9V2HHxJoss47IUrboW1raBGV0hU25BiL+FEFjYaQSgW4gkj5EGg5+WRveBt1ZSfprTFijsRvYr8hklAHyrSxwsZ+8aRTbgFk19TQ7zD6ZY5mHXUD5mgXuiBcMl+R7X50oxSg3Mqnlpbh4ACt14FAZkYXtcq1/XLeqSbDFioUse4n8xQZAnUtfuW4qstj75F+qVv30Y03ZHQsCfoKp5190RO19broPmKDJu20soPgDr60SxkgqyIw55tCPSiedV4YeUjqG/ktAMUpv91MPBv8ASgto0BFo3H7Lm30qsh+HPfpnYfnVe0A8N4DbS5Y+VlFNGInsMqx3tz31/ME0C9zL7xMotw7QNDlf/iN+8advsZYNJh48hFybuOHjpSfbcPrvVIIOhVgdPC9Bd2vZpAw/WN/zqUPtkBXsliLnXKG+RN6lBgbaOzpiAGiVlPBYp1vw5BvypqDfqHw88L8FGSNkGvXen865zRq8wM+HiYjUboqrMbaXG7/OnQohKK6YmFRcA4do7DvylAaDotNisKFVsHjJha59kRWGb9rMa5mJ2rjznEMM+HF7FpEDPfpctl+VdSFMIB/vmKcnis1yPA9j86HFezzqFYK5Q2jKZxlXwA1oOLHtLFocxGcntMd2L6ctFA9DWr2raEmV0TERKxBG8mBjseoaI6Vf9nqQbMb8gYmA9TamQwyYbjuCtr2zv/oKDZnlBQGSFtB24pcPc89QQrj+GltjFuN6kkRsVJy74ePYJpWbOylRDGRpmXeA+bXJq/Z5QbopJGpZWlIN/EUCWxWPD5mjhlRicrRZYnYA6XDjj5VGn90tLiYTfNa0UkS37otPlWhPa1DI0BZA1yrxZoyP/lBpiZnBVcJPFlXstGmUL5CwoOe0T23rTDEJx7BaMW59m1qcvsE2UxlANAQZ2Q35gixWtisDlSR5r/CEjAPmRJQvhkv2A5W9zvoje51uCtzQBKmFIuWxaKoAKoisgtzJQn1tVmbZTZTHLLFIBlWSITK+muYsVA+VPQAAZSFI5lZC3kSpFUyYZy29iEjk2DgqAT36A0ArFhcpk9rnmkZyVDSIxI6Zct/WiiixEh1Z1PFI0yRr0uQiCqjw+F1LYcacMuQjTrmNUHKOQiFeXYGU/wCGgkrxyR7t519oBsqYm1lHcV0v4rXC2vsf+2cLHBM+6bDO8mEkQBikkgAIaMKLqbLfXl6+j32LCgMFItcZl4jvDClh4iSd0mYfqM352oPkGO2JtnZ8aNjMKqRzTDDQsJYysswW4K2Oa3iBXKbML5Wy2LL2dQdcpr2n222rDNtfZ+BURhNm4dXZUyhTiMQN63Z/ZyCvI5CUjyDt3IvprmOnGg6cokww2JNh+wcDh4C4BOcsje0q2nXNXV+0cWEwwhXC75sLj8XNjsBEikC+MSDPHGRcnUWFceCVxGYXYEtLBMRxawh3TDTlXvfs/gIsRDg58TCxGy5JG2XqTljxEaMJGD3Nxrk4Wv6B2MB/Z0WAwOCxSYpVw2Fhid/du0agF7MfyrTmwBW0WMR4z7oPZI8d6CKPNNoArMvUhuPnTQ19JYkPiM352oAjTEOBu37A0BVomBzda0ph8TH2SEZTZrBAL+TaVmKYchSFw625KQp9L08TMoUiaXTRVBQj5vQG0GElB30O6YA8Sp49SDasrbPwbKTHYm1rl0Av0sda0mSV7tvrfuC/nqRSznkujnPHzY7sC/Qgig5smBjiN5DBqSCFJa3nY1SYfDqCFAbjZhz9Bet64eHVVkiHVc4+ii1WsKLcBYif1ALnxvQZIoyCoaKNgDrZ5I9PO9bRFhnyu6JmQ399tNNOQoGst8yhV6iQpr4XqlaF/wC8j8HYH5kUBuuKFyGhEdiF1II/iNvnWe2IJJkeJriw9y/0NbFfDICGjB/YkVjfra1WFwrNdIJNNTmQn6NQZk9qJT7wAIb6W08eyBUkEzN2Sh66R2PiLVtAtbKZEHTK6/O5pbS7n3s1+GpLA+tBn3igWyBetk0v1teooia5ZswPAWUZT4G9aVZHBbODf4dPpcVMsdtNG5gx5h/FegSY4lBI3RP62gPkBWdkU2yrCpvrz+orY1hluE49AD6G4qn3QFwoPDXOgXxtYGgzRqouDBCTwDLJlIPfT92MpJUDrupZST462qg0h0iWK+mtyfLgahldWCyQRE8DqvDuFqBhWLLbtAaWN2Plcrf50CxOxOUyFQORY2+Yp+fAkKGjGbhZUS49WFLkihAuBIgbmHCm3he3zoKaMuVDNK+XgHAceWa9XughvuUv1HY8rChy4fm07crkR3/6lT2bMCyxuR1Nm+QY0BZWNmZAq3tbifnepIkvuqzNrwGlvQgUkri4nG7SQDuub/ukUTS4skBt6ORDJe3hagTuwpOaNr9SD+RowEsL3A4WuwYepqZ8Smp7fcN6D6VYknBuIs1+N8/zoKyOPdzheQ3hHz4UQXtcWBAvqyOL9+g+tMWaNbExuDzCkEactVpLyOzHKiW5WRlJ7jlNqBl3K5TkYcQClvmpNAF01BA4HKSQfEH+dRIZmv2DccCGYf5lIoxC63DxXJBN309CLUAoWTNmzZOgDL/Oq0zXQIotbtoDoepAFQwx/GSvMC4+VjTFWHLdcoPC+Yn0BagzthoWIP3FzxADKPW9StaolirGHKeJKqD46G9Sg4rwBWPZmfho5I/z2omWNQuRJrm2ZY1icgdetcaeHGgZpV2hkBC23Mri3E3LNf5VrwuBSePf4d8SGLMpMKqjqe9QaDaksC3EhkAHJ3jjOp/WIpvtDMoUJiTZT2Y5YpDlv+ozVhI2nhM8a4mWY8GTEYZ3caZtHZCtqyAvHIrGNRLmz5nZVUE9FdLUHTklwgk+8GNQ3FhIEAv+81qKFsOXJByqFzdoxk38FuKDENNNBafZxUqgyuk5LKBzzG6VgVMM66B95wIBgk08QwPyoOmz4DPY4hUbiAseUeqqBTRLhF4bTFh2gLyHXpda5SB41UnD4oHVUkWNrm3ICS4pgacoCrwsAf0c0aRym+mlhr60HTDs4DIzTXPZZpyq+PEn5UqWTbCDO2GQp1RrDT9cka1lSfGrYLhVykEBY94HBvyZTeiMuKa98MWKtcCRpGtfXnzoATH45yEIjc/CDJl07iTaujmIjBlindj/AHaBioI0P3hFc14GkLf7HiEcm5sxRF7u1GajREXDJiHbg+5Rgpt0exv6UGh8YBpHGAP1sj/IAUyLEo4Nlw6svuiWOTMx7slYjh0AJEGJjHw3Y5vVo6auQJYxTXylTIMhax5ElKDSMbiUGuH7XLKrAW/eNWMbiAMxypwtqTqeRFwaQshyCMTDLr2rhCbi3Ek/Sh9mD3aOY8bEGRPpQbM88ihiXYWvZHc6c9DrS455A9hJPa3A52t4qRalpG8ZIK51PvKVzXPfYEfOqkSIEDdSI445WAHo1A7Ex4XGo2HxWFgxMcgtlkwqyXF+BLajyOlfKtu4GLZu18fhIFywQTxvCruXyqyJIFLXJPHrX1SKIC2VJ72vcMqPr+yb18++2kW727M2WRd9hMHKBKbk/dbsm/lQednabCoJMuUyq2R3ByMo1ut6+q7Ij9l2VgIjhp4ZDGsslpTLrIL8XQ+XaNuFzavnO1XXFbI+zEoN8QntezHv+HDNGY9fBxX1CCNvZMLuzMIxh4RGEFhk3YI4n8qBiTIrDOXNgf0g0tyBsL0bSvJl3LQm1zlRZAV/nS0ixjWIznoxIHmSTehkjmDDelwb6k3PoSLUBN7WbEyaE2XKANfO1WmIx8fZMWbTUuuYjwKGiQJYgTKL8RLbx0FNWOLK1zDckWKOAxPda4oM5xUrC2SdST/dMwuT3MaC2JNyJJFI5OHU+oBFGyRqzXlmJBFmDEsCOd9KEyhmvJmHegAPnragJv7URFf7104Ax5Xtbrk1okxO0tL76xNu1Fb/ALauGHDSHRyrdZMoUc/ezU9sViIxlzqykW0zWAHQkUAjHMoAlUNrrlJQ26G2lG2JwroxWTERnWy2hkQHjq5F6ztJmYFYxmv2vutLnvB1o1ZXYpPHudLAqMnnaRh9aC4opMR2lMLjheygk+FKMiQuyPDdlNjwtcctCK1ez4ZbNHJI4BF90SQPHKTVFcICS0ZbX+9il9NRagi49MlkwmFv0YNe3hnvS97nLHsR8Du1U2v4E0QXBEHtqua+hiUkeDFQaHc4e4yTup17RDgG3TSgasroRqDckgEXuemqn61DKysSEQcOcg8xe1MVHVbJi4z45HBvyswBrMVVGLIYnB0ayqbf/rN6BryTlQV3chsdcxJ07ib1n9on1YxCw1P3ZHzNqMTyA3EkJ04GN1tyte1GZWlBBMWtuDEcOhJtQAcaTltDOoGvZjZ/mDRjGtbNmluOKmNgCPEijggw4N2kkLKbgCVDc+VBLhnLFoUks2pvJbXyJoGNj1EYLYV78maQ6+HZ/Ole3qSM6ZRr7qZiP4noQmMQggPcCyhjnF+oN+NE0GOcDMjs3PMqmgNZcIRmstzxvDKPP3jRLNA36OQhr6AqVXh0c1n3eI0VlDD8Jvbw0YULoALezwp+w7sRb9pzQOlmxMZ7UkgBOgu6DysSKCOZmLA579xDafvClII1PaIPUEkD61pRY2U5GCluNt5r3DKTQMQyi2WVlzWC5kPzIuKuSKdwWaaQMDxMKtf94NeoqYpb3MhSwsA2ZfNZAajQhjdoJEPEHDqYz6KFFBjIlRiTNIbcRu5B/lNF7ay2DLGV5Zkk+vGtZjDW3hmJ4hiIDID3nQ0toYx2Hd9BwO78eIagobSjPEIBwtG8i+YDil+1Z7hWk8RMQRQmHDN8UY16tf0Bq/ZsKbZpEtyALA+pF6CzKzFVRsQ573DWPmasy4sAjNLY+6GRGHzv9apsLhAukEjE+6wle9+5ctqSsMKksGZSOJvmt6AUBmPEWzMIlHVkQX/hFSiIVdBi2OnxBgv+b8qlBzsRITYnfg/ila97chlB+tAkmJOaODE4SPQEiRnjc+ZF6zrdlIOMhi1vZoHbzBAP0oW9sYOqYoyKtiSCYo7D/mAUDW9pZCGhDgXs0Rmbh0u35UqJWz3Lywka3dJFJ8LCkpLIWUNiZAL8S5sPIX+ldeLE4YDKdoLn4FS08yk8eIS3yoMq4vExaST4kqLAZo7jTvuDTIsY80lpHQxjtHfxREW7vio5ZMKb2aC34hh5wT5EZflVLNhVyhoIHI5vAIzbrcfyoGTSYaTKqzynuTOFHirXoFUm/bPZ1zsQbDuUgU9Z2B3kbRpGdLK5IYnW1wQb+VG0k5UBsM8q6m8bYjNr32tQAphZR96pl+EsqAEd40b509UwjBc28Z/i1BXToAxpaEWbKkt+aPLx/jjvUVpbsVggzCxQsQ/oSMtBoMWHkBEZJZRYKkhQDnfTSs0qYnJc4hjYj7tXUHXn2bUx58SFuMJBGPikTdX/AMVz8qyJIwY3IW97tMAbX5XCXoI82MBCNLiBy+8YnTwuaYjkC4xMtx8ISUX9TajjaSRt3GzHmd3IgzgdNKuRVRsssYzaG7ytex56UC3kaUjeyQMc3CWJQxt35fzoliRRYQRlzwKzDXuy5h9KcrLGuVFXtWu0RYkDvJBoWgwZUyB5WkucyuGKkdbhL0ED4iEXXAKOWa8h06XU0pmnYkmAdo8MrUZihAveMG18q/kf9KCzrwLgW+FXIPrQGsmJWyiErw4bz6E2rxn29SR8TsXFtfM+Emwzk31aGXeAa9z17BGkuSGkB0uLaeh/lXL+2Ecs/wBnnOZ2GCxeGnsUXg94Dc2v8QoPBDDwtszD4psReCLbGHE6AOHw7SQsSHHRsiWI/LT6gk4Iid4IiciHL2xYkDTQ8q+W7LUTSY3BMAUxeHikII7RfC4iPEAA+GcedfWJAjSMTDftHVZMxPo1qAUmiGrKhvbgNRbutR76Eg/eOobSwUAeVlqiirqI5l/a4eWt6IyR5SGiI00K5B6/F86BaZGfKXKjq65vQBRRkFWsJIWB5FTfzAWqIQgFDdjoVIdSB43p0eExjWYGNVtozuVsPEi9BSmRTZYlZgLDJFK3mMy2odbgOVU8DbIG9CBQToYSoMkcjak7lmZR3X0+tZ9Cb5GHr+ZoNiQhycjHu7Ka+YajaKWOxsypw94AX8TpWJbpqtweOhINECWNzErHUkkX8zQb17zNbQ3jyP65GB+VQiTIArixOlwv0kBFZVIQj7lTfU5Q4t5i1WJoQQWgcnlaaUAeAcmgafbSRurCw/u9xfx0oWTG8JIS5H440OnjSmmhY6LiU1uArqbeHZBohMhBAxOPGvxKrDz7YoBKLezRKjDXUaHwC0UawZxvAhTW6hWBFEyOfdxczDoqSr6gOardYq+m+JJIBJKg/wAZFApxh1uMoseBKC4/hNQJCwBjTXgbSLcnrlvRMuKQ6kix1zEkaUW+ZiBMI7dVXKT/AAaUDQkYAjy40P8AqsCp/wAVWEK6NiDEw1USGUfMEigL+60S2F7XMpkv5MlKNnYlpLEk6NHYX8rUBsDcAtBIOZRiy/41pzSMVCIqIByWSEW/wis4WS/vwgczlv8AkTRmJrfdyoAQLgqdT4EGgJQzaszd4yhuHeCDTBFKe0MyhuBSPN6C9ZSZk0Z4NDyVr/5RQHE4q9xKARpdUUH5AUG/2eVxYGPuZ4srW7yL0lsNKpOWRT/8bfXLSTi8Uffldja3aiXhVw4g3JklnQA6ZFNvkaCzhpCQQy5ugSUH5LRJhptDkj/jKMPJrVoWRWAyY0tb4MQHA8uH1pM6S8VeEnnkAP50DHLgZL8eIbEl/pSgzLa27v03rflUWSa1jkewPGNRkHij/WhWV7X3aMb6hgwA81agaJZhewg/jLfJqFp8UdA8YBtYAWHnYUsvKP0kNl4XWx17s4oCJjbKARyzoB5XU0BscQ5Abcg69pQ9z5qaE9A4Bve4aT6E1RXFKNUGvDUfIUPtOW2bFxqw0y5AxHcclARjY/Gt+d2J+RNUsM6A3KWI4SX18BRLidGK7RhVhqQYQl/WiXFoykNtHDs9+yDA7N65aABHIMufcgHUXVzf0NShkKk5t5hH0Qs0UZGrjNlNl8alBx2mjzANPh2Ua3IlZfmppyywTNlOJwyhR/eYcZfI8flXNMUAuQZNLaMw1/hoo4sxAVc1+K52BoOlnjYMm9woueMMccbHvzFLU+yuhjVXkBAvl2ioFve4KAtY4o5YDfdiwtplEjWPDWmzTYSYDeRMhBGjMVN/C1AxMNJIwHss+QaJnmk3Y8HyZaCaAR3DCGMk3BOJZr93aP5VcWKwsTFo2yG3N8R6Axip7biJyQMVKpTVUjErHzJoKVYhlMaYIMRozTxv6iUGtC4nFBN2s2FAvoMi7u3cFirny7RwqBxNtbCow4pPiMLnv4MxaufiNv7Dw6rvtpRScsuEQTSee7AHqaD0sYnylt/g7kXKRwzPIR0uqihefEgFcgW3wiJwPnJpXn4ftT9l1Ts7Ux0Lnk+Dmsg56Jz86p/tF9mMzH+1TINDc4XF5vQpQd5DKx7akcjYcfV7UwhU1Vn04hkbXzQ15yT7T/ZpFsmPZ7a2jw2Kvw5ZkArbgdrbNxvawW0sJm0+7xE64ebXgN3Plv5XoOzFMWtmZVtrd45df4DemGWdf0aFhyKowUjqLtmrK3tzWzhGI6bsafu2NNglxCA6ELr2Rc/Sgek07K28CoBYEETqwB59ljS5RJowxKfqhpJlNu7e/wA6o76QqyxuxHDds17+GaoWlcZLYw3axUsGHhlNAUM8rALJIjdBK9wB0sdKe0eDYAsEQn4oXy6+BJFefxe39gYHMmIx0QmVzGYcOgxMikaHeiM6VysT9t8Bhw42bDLiZdSGnX2eBT1K/pD/AIaD2EmCAXeQSyyKqmSRzYIgHN2NlrkbWxGzZ9ibew6Y/DyS+xSOESVWuY2WQAFTbl1r57tH7SfaPaYdMbj5pIXZT7Oh3eGXLwCwx2X5GkjFMMPMl3DMgRgwsASdQLUGjYBA21skkXDTFSO4owr6naSLtvINT7qzKJB/Bf618r2DJAu2NjNM+7j9sjjdn0AEl4wb+JFb4tv7TwWKljkxDTRxzyJJFiXZgQrEEBmNx5UH0csjABZ2QnjneQ3/AIdPlULtEVZJ4SbWJyPw/fBFYdn477NbRjU4fGMJzlHs2JKRSqx5Ajj5CtsmECNZjIt/dziVR/jWgVcWY5oySddD8rCjjaYG0bMT0XX5VW5y6bwm1uFzf/CDWiJY2IupJ/VYoPPT86A0lMd957NrxEkb5v8AAKS0mDJGYNf/AO0wA9GWtDLKFCoZkCk+7KoHqxvWc4PeG+ZRm17U2HY37xnvQCHwp/4w7rqw+Qol9iJ7TzL4Jf60PsSi5WZyOgVvy0qDCOtsxkFuJa4FvQ0GlYsGxCxT4l7DiY7fnS2R4vdsRxBZ0/1pyQK0ZizCUnlnU7vv1jvS3wcCXGeQAC/ZeKUH+ApQKMoA7aqB4D8qsTYUakXOvZsbHzDUITBkqN5KCOBmTT/CxPyoTa57COOoa3pcCg0JJs97ZjIDztaw/rxqzJhlP3WMdRa4DKw18QTWYImUu0cZ4C29jzDyU3oWhhtcFNeQY3FBrE7MyD28gHiWJKr52oC1nKiTDt0cEMD5WvWZMKXP3eZj0RkY+l70pgikqXcEG1iBcHhyNBuEs4uUlw1+ecWI6aMtWsu0nsgOF1vosata2vJKyJGzFVWZdTYBrLr+9TtziEBzSKADrpmv3gDX5UE3kl8rZA97EbnJb1NERiBmAgWQDVTGjFyPANaltFI1i0nHS7Ryga/tilNhSpNgjW5gadeNBZbGX7OFbuG8ZCf3XNAuI2oGOWAKwOUiwJB6G73qBT7uUBuRBYGrs6i+aRCdL2JBoHpiNoiwkBBOhJjGgPA8TS3OIzEvGwFzZmhuptzFjRqMQLMJJWC8THddByuRVS4jELldM0X67FU637bf6UC23yqFYIytYiyvHx8P50sxyLdYcK4DayLvGIbhxzC/Smb3FOdCrFtOwVYkeWtDmYaFdR1zA/SgkeFxJVvuY1sbsMrm47yKI4YXXsKTcm++OW3QA2+tHAd44W7kE5SFJOvTTWuftf7Q7H2XnUYlsZjIuGEjNyr8Ms0y3RQPX9Wg3buNQ14WVjqcj5lv0ysbUxJCGGVYkZuLMiW868HiPtvt2Un2aDAYUcPu4nmNvGdyv+Giw/212xED7Xh8FitCNU3DesfZ/wAFB7cnEXjIlQgdhd8gMYJ56G4oGMgujxwup0uY5rac9DXAw/202G+59owGLge13eERSoH52F1PyrtYbbf2fxuUYfaaK8jG6Yi0Uqnj7kuQfOgvcwsGDll/DkU2W/cyk/Oq9niGXdlMQl23mZt2QSLAWCWuPGutFhsS7JJGwlULdd3G3b00PZa3zqNHOLmyuDc2kOWxPQ5T9aDiez4lLH2fDWtbMpCsRw95NL99Sus2GkIs6IHJ5TKqnvs9Sg42KmweihYSVN1W+dde8R/nSYsXCNFw+HUc83tKA+cbD6Vu3MGISR0jgCxxl5Jmw8MEYC9GaTjXNyG91hsPxlW17tLrQdCGczEhYAjajTEuBbr2g9BIMMQc0UDtc9tMSJMp49rIgNXFLhgv3gw4lA0QuyNl/FpGB86JpApcCGNHBvlSCJzYi+pcGgzR4ctly4ZHBsXMe9uqjUm5Ir5rtbbeL2jiMQYpJYMAXZcPhkkYKIgbDeW4nrevoe2Nsf2VsraMq5d5LG+Dwqskaky4hClwAoOgzNxr5QqgWA5ACgsLboKLKxtc+F9BUAGtWDb/AF4UFqjN8S386IrInEAjjddQfAihBtyHp/Km71bKCDwIIHDxNAosTe9X2iNQD5CmMUZRcNfiDyt011qgq8iPOg7WyftFtjZy+zxyiWAKd3FiryRw21vFcgjwvaik+1W25WkLYl1LtmCwAQxpbkojtp1veuJqBoV6aE1QQntWNr2JHC9B1sbtTamJEch21iZCy9qNHkhAJ4jdxKkYtw50pNtfaFI90m1doKgXKAMTLoOnG9Y1JCEW0t0F7c9aIcCRzHdQKijWRwhdUBvd3uQD321rrRbIm3RY+z72znKzuYyEF7iULu/nXOUDhYWPhaiXFTrGcMqoseZScqgFso0uRQLeJjI4VQLEg2NxcaaE1eSVBcKNNSVYkjlwpoJ6C/yoQSCf6tQEs+MjAkVLFCrhwqZlKkMDcilIk0pY2Ba5JJ4sSbm560/eAi5HEWa3MHTWoLpcRWKnKbtlzA9LmgoB1YLns3MlTcEd97104dq/aDBsk0O0MX2SCQ880sDBbdmRZCePOudvpFIDAkg/EAedOZxPC0AADFlKWspv0vQe0wn2ywuOmGGklmwkYjLNNMVeINGCSv3Yz26Gih+1P2flkSJsZiEVuEs0Eqx+qkv/AIa+fqigqUHaIs+pBDHskBT/AFrS1U2IIsQxHAjnyoPff+tNjR4gxxptMxZspxItHbv3OYvbzr0UO0IMRhRjFxiDClcxxjytHGq8Lu+bjyItXx6N2RjlcC7X0AJB6aimytPLGsJnlMSyGVYnY7oOdC4QaX8qD3G1ftquFAw2yJo8UwJ3s88QkwoH4IlYAt48K4c32v8AtLijDFDPhsEODnBRLEJGJOrtJfuHIfWvN7tl97rbx8DRmwsANbUG5trbdZysm08exB1viZRrwPBrVsi219okjRYtp41Qb6TOJUNtdTIDXIjAUny6Xv3U4NbraxFxfS/Sg9ThPtrtKJoDj8HhpcOWySPEjxzNYcUYuY7+VepwP2m2Dj23eGkXfzXWPDzoUnzAcLghT3WavmL7p0uztvgTIvFopQbXBA1B7/6ORWaNi6WuGDI4vmjdGDK6njcWoPsntkpGV4MLIo4b2IuR3ZmN/nSHkmlPYhw0fL7uMf8AdembI2hDtHZsG08U0OBLtMMRHiJmiYGC+8cKxvbS40rx2O+2+2cXKybPOHwWFErrBIkQfEyRXsrvLOCQbWJsooPaxJtSAqAqEMurPhs3YPJW0ot3IVBfDpHGTYSCJmBPDkTXyg4/aqyzMcfi95KWEznEzfeG/Mg603CbS2rgp99FjMUHtYETyMPBgSQRy1oPqvs5WxeBCFPFdL26i5pyS4ZQLxEaaqixju97SvnEv2x+0BkYw4pY21QgYeJkNvjBmzvfzrbhPtztQLEm0cLhcZEr/eSor4ecx9LRERk+K0HumGAmGuYG9lUyMR56mlNg8OCRvZjpcIkJvfpcsPpVbOxWwtqRQzRTYOPfRLKISEXERBvhlWVzr4GuZtL7RfZrZUrRPjXxGIjIVoMBGJSpHas0zlYvS/yoNrQPnCrG4Y8ndSTT3wkscJxM2eCCONpppXKZI40F3Zjmr59j/txt/EOyYKVMJhFuIo1SKaUj8UksinXwArgS4zHTqRiMXiZixJO+mkk468GNB29pfa3bWMabD4LES4PZ4ciNcP8Ad4iZVNg80vv68bA/zrhzS4jEm8+IxE+gF8RI0hsNbXYmlAi+nAaf6mnLY24kX14cO6gBVCkFSy2OhU2It0tXSj2vtrDKI49pYxVsLhpGceN3v8qwBowxDWA5XF7+lJllRntEpRcoBF73I4nhzoOrL9ovtDJFJhztCdkkBRmOXeFCLZRIBnA865IUAcqoG1Xe+l9BQNVVAuQDcaE8qNhFlzWBIHppSM5tUaQZbcz/AFagGx4+N/OiGU6BbjnbnSyb86YrAA6206UGyHE7QgUrh8ZioEPwwzyoDy4KwFE+0NrSEvJtHGsVFs74mYkAaAXLVi3wso1uB5UTSRFe7kDpfvIoNkW1NsxFRHtLHKEOdQuIlsCe4m1SsQcdkKNdb36VKD7hAdhQ5FwsuFC4gscQsTJJwG7VrrdR6Vx22AhnIwePGTIJIHlY53k6MEAt3GvjGFxc+DmSbCs8UqXXPG2Rip4g17bB7R2m+Dw7zYnE3eMhRO4eUJnN80igXB5UHq3mwmFB9qkjxITeRFbZkcjQG+h86zL9qfshh8QRi5EinR2jKrDLMsVxltJIi2PQ/wBX8JtjaOJgWJY5jvZs3aUk7uIaWF+ZrzLMSdSaD1H2v+0WG29i8JHgY5ItmYBJRh1lAEss0xvLK+p42UKM2gHfYeYvYmoBa9VzoGZr2HK1UD8qq2v9cauw491BdyasX0qqscxb/SgNWZSdONQknXhVDhc1d9DQUL0xWtoeBoASRyqHNpcDTprQNVrkjXS4NXdiGbTrYafTSkq1if8AxTAxF+pvQQOLaX9aJVNyx7jytSAb+tMzZR+0bHnQMd2HnzoQWcgfD8RqgeZ6aDpV3J4XoHFhbLwGvDifE0BY6EGx5UFhpcknu0FORYjqcoPjf1vQUrkaN1uQdbmrGRuIv56+dEAOIEZ1HvDr0vTTHLcZhCoOUg8dO6woA+7OReGvvDVvM9Kcpmsw3gYQq1kJJsp1bLfSls2QlWSO4PAfkaOOUIQ4AzA6WAB16HjQAuQMBkXMdBmUX49TWnDxw7w7xBazA+8ACdM11N9OPOhBja/ZZbfi7S0ccEpBy20Aa4kTtD9XNQXiIIYnMRZFV1OWdQXADE6snkOh1rJJgJ0jjmMuGKyBGAR3Z9eItktpz1rpLhtr43C7tMNjXigbeXw2GzK7KuXPICc/ZGnT1rI8GOw0RWaUxxdkLDvQcznU9hSRegVHgMW4UxbuS4Nljdc2h1IDWPpSxHPEzQurJMjMrIdDca8K0JFiHYrCVmuqt92QxbUDQMAb6ivV4GTHbWw/s7YeF8Rhojh8bniSDES5TYATqd6Tbjbqb6DUPIbmQaMtjxtcXI8BSHC3UK66rquoPE6EnSvcYn7PbR3eKXB4FI1S25XFSxMFcmxSLKA4FtRy0rzkuLxMU7Jjmw2LeFpomiQRsmcJu1IeNLWHjrQc98Zj54dy8gZQCg7ChzGDfIWHEVUOZVFl1NvfB8eFQzXa+SOx0CoiJGLacFFaDN7QyOxiVwqq1zlDlfiYnT+u/UEuzgi3QHqKKMvcFQp090oDe2tViJWV0zRBLgOOBVh+IWNqVvJH1JNiScq6AX6AUDpEiYAgkMb3CXJA43AP86XdluqtcactLd4NN9olWMKcoVFyhlHa1Yt73H/x3Vm7RBLFjyuTrYcBQOWVg8ZLq4iuFV0DplPFQDpTGghldt1aNSLDem7+JMYAv5VkVuOVbsABoNPM06LNcM7XbuuAPCgtsNiIgoMV0JNiLMDbjZhSigPuXF+CnW3cDXQR57WQgWJJsx1vx0tSZYmftB3sWJa4UgHr1oMqxzEMQpawuSuungKtVnYLlU5TwPAG2lLLvHNbMgCDUx2XQaXsvOtCviJt2IlCIgFri4PrQC0M9hcpfhbNqL87Uizo0iuLMCBrxrqBbc7gCzXAJNBPhBLGzhcropZMpvmF75SOvSg51xrrVZgPA0OvpUvwNAdz0qWJ1vQg1YPG3lQXbvofOpcnTlzPXuqyL8BQUCCdelXbpVaXHfTBYC96CAEW5VKmbmfKpQc+NluS9yrCxC2vfkRennFSNh1gzGySBky6C3Qg9+tJxX+8Yj/nS/5jQj35PE/WgjM7tnlZmb9Y3PzqlAJ7uNRuIq14+RoLNUON+lXzHh+VUvPwoLv/ADqC1V1oqCADnV3GvGoeJqhQEAx90E+RqyCo5X5i3Wnw8PI0qX3z4D6UAacjbx5VRJtofSofeNXyHj+dBM7f+ajM1hw4iqPPxqN+YoCJA5D+Zo76agX5UHSiP8qCAi9zy6/WiBvck6Uvp4D60zrQWSOA51BQj3l86I8D+yPrQNUE3ubDvoyzCMqG008fKgPu+tVyPj+dA5WNlVrMgbNZuRI5VojTCTDLnEUi6k3urWNtAevjWMcfI1Of71B2Hwwy3WVN4qGyxKQjgAe8eR6XGuvdmzrvFinAUSxrYCyq2Ukg9tW5d9XD+lb93/LQn/eJ/P8AzUDI8Q/tDouChcREq0c/+0FI7i6Asc9vCtk+2ZJY5sKmB2bhYi5QeyYDDGVQBYRmSZTJYetc/Zv+9Q/+5j/6yUl/0uJ/5sv+Y0HSDbMvHJJI0Ezzl55YoXlR4XRTaNEewYG+vhwy67Hx2ChxsLYX2wQx4R+ziZMNit9xGVS0boFHEAg+VYNo/othf/iYfq9Uf0GH/wDm/wA0VAeLxWKmmJeZksIxKoxUkTyAgfAHK28BburG2EiciTK8cJZRdGDEEnUFmAXw1ruv+gxH/tW/6a1gxP8AuGE/Yw3/AE6DmNhh96UdxGvuNJGAc1zYGzc+tAIcUY2kVQypocpFxpfh602b4P2B9aBP0M3iv1oMGdzYMTYHQHQDXpTBLltwt0tWnF+7H/y4voK57e/50GrM5Uk3txt1oc+YBm62sf8ASiPClD9H+9+dBqhIdMx95TlHQAcOFP0PaFib3a9zfWsmF9yTxNak5+FA1Zt2GYuVAB52FjpbSseI2hIwKRdlNdeZqYr9F+8KwN+VBtwyII2kfKSbe9awHHnWwSpa41vWI/ooP3adL7jUGgSo3u2Nrj+r0wPGI5HbEWmVk3UAjz50PFjIH0t0sa5Pxjw/Kmp/eeJoDlhaR2ZCvaJYgnLrzsTWaZJY2KSIyNobMLGx1B8K6C8R+zQ7V/SYL/2cf+Z6DAgJozmFh14VIvdWjHvr+waCgFt4VRB40Y4VR4LQA3C9qoXPGraiXgtANiTUoxw9alB//9k=",
            "icon_url": "https://pic1.zhimg.com/80/v2-a9cb8a0bcebdc99eec26fe1dab885e83_1440w.webp?source=1940ef5c",
             "gestbooks": [ {
              admin_id:0,
               book_id:1,
               create_time:'114514',
               user_id:114,
               avatar_url: 'https://tse3-mm.cn.bing.net/th/id/OIP-C.5zVqqn_z2vFdWoeyeBPyiwHaHa?w=184&h=184&c=7&r=0&o=5&dpr=1.5&pid=1.7',
               user_name: '路过二刺螈1',
               context: '别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师别卷了吴老师',
             },
               {
                 admin_id:0,
                 book_id:1,
                 create_time:'114514',
                 user_id:114,
                 avatar_url: 'https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0',
                 user_name: '路过二刺螈2',
                 context: '别卷了吴老师',
               },
               {
                 admin_id:0,
                 book_id:1,
                 create_time:'114514',
                 user_id:114,
                 avatar_url: 'https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0',
                 user_name: '路过二刺螈2',
                 context: '别卷了吴老师',
               },
               {
                 admin_id:0,
                 book_id:1,
                 create_time:'114514',
                 user_id:114,
                 avatar_url: 'https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0',
                 user_name: '路过二刺螈2',
                 context: '别卷了吴老师',
               },
               {
                 admin_id:0,
                 book_id:1,
                 create_time:'114514',
                 user_id:114,
                 avatar_url: 'https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0',
                 user_name: '路过二刺螈2',
                 context: '别卷了吴老师',
               },],
          },
          {
            "admin_id": 2,
            "title": "妈妈",
            "signature": "哼哼哼啊啊啊啊啊啊啊啊啊啊啊",
            "image_url": "https://pic1.zhimg.com/80/v2-0ddc39b65517f14c07fbcd3f8b0f8cb8_1440w.webp",
            "icon_url": "https://ts1.cn.mm.bing.net/th/id/R-C.1efb8caf30b6518b35bda6212687ccd6?rik=GUmOOnexRSG3Lg&riu=http%3a%2f%2fch9989.com%2fuploads%2fallimg%2f220716%2f0930294157_0.jpg&ehk=MTqLvCPczteKJZRYkNIrdUq4FUy8JYwzwgR1rZ7b3Do%3d&risl=&pid=ImgRaw&r=0",
            "gestbooks": [],
          },
          {
            "admin_id": 3,
            "title": "yorushika死忠粉",
            "signature": "哼哼哼啊啊啊啊啊啊啊啊啊啊啊",
            "image_url": "https://wallpaperaccess.com/full/5679257.png",
            "icon_url": "https://tse3-mm.cn.bing.net/th/id/OIP-C.5zVqqn_z2vFdWoeyeBPyiwHaHa?w=184&h=184&c=7&r=0&o=5&dpr=1.5&pid=1.7",
            "gestbooks": [],
          },

        ];
        return p as Admin[];
      },
    },
      data() {
       return {
         admins: [] as Admin[],
         table1:false,
         table2:false,
         table3:false,
         dialogVisible:false,

      };
    },
     async created() {
       this.admins = await this.getcard();
       return {admins: [] as Admin[],

       }
     },
  };


  </script>

  <template>
    <div v-if="admins.length">
      <header class="header">
        <div class="header-text">
        <h1 class="heading-primary">

            <span class="heading-primary-main">菜 鸟 营</span>
            <span class="header-primary-sub">Rookiable Coven</span>
        </h1>
        </div>
        <div class="header-arrow" @click="hitarrow"> </div>
      </header>
      <el-dialog
          v-model="dialogVisible"
          :title="'cin>>'+whoguest+'>>endl;'"
          width="30%"
          style="background-color:rgb(258,258,258)"
      >
        <el-input  v-model="guestbookfrom.context"></el-input>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false;guestbookbut()">
          Confirm
        </el-button>
      </span>
        </template>
      </el-dialog>
    <div class="card" :style="{ 'background-image': 'url('+admins[0].image_url+')' }" >
      <div class="xbox">
      <el-image :src="admins[0].icon_url" fit="fill" :lazy="true" >
      </el-image>
      <div class="name-box" >吴老师</div>
        <div class="good-box">
          <svg  width="45" height="45" viewBox="0 0 20 21" fill="none" ><path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/></svg>
        <p>114514</p>
        </div>
        <div class="good-box">
          <svg width="45" height="45" viewBox="0 0 19 18" fill="none" :style="{cursor:'pointer'}" id="0" @click="tableview(1)"><path clip-rule="evenodd" d="M.733 2.8a2 2 0 012-2h13.2a2 2 0 012 2v6.015a6 6 0 01-6 6H5.534a4 4 0 00-2.189.652L1.507 16.67a.5.5 0 01-.774-.418V2.8z" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><path d="M4.5 5.62h5" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          <p> 114514</p>
        </div>
        </div>
      <div class="xbox">
        <div class="but-box">
        <div class="desi-box">{{admins[0].title}}</div>
          <el-button type="primary"  @click="hideButtons(1)" class="show" color="#2e5496">撅一下</el-button>
        </div>
      <div class="sing-box">{{admins[0].signature}}</div>
      </div>
      <div class="xbox" style="width: 400px">
        <el-button type="primary" class="hide" id="1-1" style="display: none" color="#2e5496"><svg width="24" height="26.666666666666664" viewBox="0 0 18 20" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M15.882 9.285c.266.33.535.69.653.94l.024.052c.263.555.5 1.054.843 2.069.354 1.046.917 2.977 0 3.31-.867.313-2.035-1.59-2.035-1.59 0 1.147-.64 2.645-2.02 3.726.51.042 1.204.215 1.609.473.405.259.405.628.183.866-.221.238-1.562.23-3.057.217-1.283-.012-2.615-.081-3.239-.146-.656.068-2.06.15-3.368.155-1.385.006-2.467 0-2.745-.226-.277-.226-.228-.593.079-.83s.993-.445 1.53-.507c-1.382-1.083-2.02-2.58-2.02-3.729 0 0-.67 1.896-1.687 1.59-.754-.227-.57-2.271-.275-3.31.24-.839.43-1.27.647-1.758l.158-.362c.107-.25.32-.556.546-.844 1.206.828 2.701 1.109 2.701 1.109l-.15 1.149a.904.904 0 001.787.264l.19-1.15c5.005.475 8.26-.769 9.646-1.468zm-.692-1.14c.003-.115.005-.237.005-.325 0-3.601-1.835-7.22-6.35-7.22C4.327.6 2.49 4.219 2.49 7.819c0 .085.002.201.004.312.68.23 6.853 2.204 12.696.014z" fill="#6FADFF"/></svg></el-button>
        <el-button type="primary" class="hide" id="1-2" style="display: none" :icon="Notebook" color="#2e5496"></el-button>
        <el-button type="primary" class="hide" id="1-3" style="display: none" :icon="Comment" color="#2e5496" @click="dialogVisible=true;whoguest=admins[0].title;guestbookfrom.user_id=1"></el-button>
        <el-button type="primary" class="hide" id="1-4" style="display: none" color="#2e5496" @click="likebut(0)"><svg width="24" height="25.2" viewBox="0 0 20 21" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/>
        </svg></el-button>
      </div>
    </div>
      <el-table
          v-if="table1"
          :data="admins[0].gestbooks"
          style="width: 860px; margin: 0 auto; display: flex; flex-direction: row;margin-left: -35px;
          border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
                   "
          :show-header="false"
          id="gb-1"
          class="table"
          max-height="250"
      >
        <el-table-column label="头像" prop="avatar">
          <template #default="scope">
            <div style="
        display: flex;
        margin-left: 20px;
        flex-direction: column;
        width: auto;
        height: auto;

      ">
              <div style="display: flex;width: auto;height: auto;" >
              <el-image :src="scope.row.avatar_url" fit="fill" :lazy="true" style="
          width: 50px;
          height: 50px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin-top: 10px;
          margin-right:10px;
          margin-bottom: 10px;
" >
              </el-image>
              <div class="name" style="font-size: 18px; width: auto; white-space: normal;display: flex;flex-direction: row">{{ scope.row.user_name }}
                <br>
                {{scope.row.create_time}}
              </div>
              </div>
              <div style="margin-left: 20px; width: auto; white-space: normal">{{ scope.row.context }}</div>
            </div>
          </template>
        </el-table-column>
      </el-table>



      <div class="card" :style="{ 'background-image': 'url('+admins[1].image_url+')' }">   <div class="xbox">
      <el-image :src="admins[1].icon_url" fit="fill" :lazy="true" >
      </el-image>
      <div class="name-box">牟老师</div>
      <div class="good-box">
        <svg  width="45" height="45" viewBox="0 0 20 21" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/></svg>
        <p>114514</p>
      </div>
      <div class="good-box">
        <svg width="45" height="45" viewBox="0 0 19 18" fill="none" :style="{cursor:'pointer'}" id="1" @click="tableview(2)"><path clip-rule="evenodd" d="M.733 2.8a2 2 0 012-2h13.2a2 2 0 012 2v6.015a6 6 0 01-6 6H5.534a4 4 0 00-2.189.652L1.507 16.67a.5.5 0 01-.774-.418V2.8z" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><path d="M4.5 5.62h5" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
        <p> 114514</p>
      </div>
    </div>
      <div class="xbox">
        <div class="but-box">
          <div class="desi-box">{{admins[1].title}}</div>
          <el-button type="primary"  @click="hideButtons(2)" class="show" color="#5f962e">撅一下</el-button>
        </div>
        <div class="sing-box">{{admins[1].signature}}</div>
      </div>
      <div class="xbox">
        <el-button type="primary" class="hide" id="2-1" style="display: none" color="#5f962e" ><svg width="24" height="26.666666666666664" viewBox="0 0 18 20" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M15.882 9.285c.266.33.535.69.653.94l.024.052c.263.555.5 1.054.843 2.069.354 1.046.917 2.977 0 3.31-.867.313-2.035-1.59-2.035-1.59 0 1.147-.64 2.645-2.02 3.726.51.042 1.204.215 1.609.473.405.259.405.628.183.866-.221.238-1.562.23-3.057.217-1.283-.012-2.615-.081-3.239-.146-.656.068-2.06.15-3.368.155-1.385.006-2.467 0-2.745-.226-.277-.226-.228-.593.079-.83s.993-.445 1.53-.507c-1.382-1.083-2.02-2.58-2.02-3.729 0 0-.67 1.896-1.687 1.59-.754-.227-.57-2.271-.275-3.31.24-.839.43-1.27.647-1.758l.158-.362c.107-.25.32-.556.546-.844 1.206.828 2.701 1.109 2.701 1.109l-.15 1.149a.904.904 0 001.787.264l.19-1.15c5.005.475 8.26-.769 9.646-1.468zm-.692-1.14c.003-.115.005-.237.005-.325 0-3.601-1.835-7.22-6.35-7.22C4.327.6 2.49 4.219 2.49 7.819c0 .085.002.201.004.312.68.23 6.853 2.204 12.696.014z" fill="#6FADFF"/></svg></el-button>
        <el-button type="primary" class="hide" id="2-2" style="display: none" :icon="Notebook" color="#5f962e"></el-button>
        <el-button type="primary" class="hide" id="2-3" style="display: none" :icon="Comment" color="#5f962e" @click="dialogVisible=true;whoguest=admins[1].title;guestbookfrom.user_id=2"></el-button>
        <el-button type="primary" class="hide" id="2-4" style="display: none" color="#5f962e" @click="likebut(1)"><svg width="24" height="25.2" viewBox="0 0 20 21" fill="none">
          <path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/>
        </svg></el-button>

      </div></div>
      <el-table
          v-if="table2"
          :data="admins[1].gestbooks"
          style="width: 860px; margin: 0 auto; display: flex; flex-direction: row;margin-left: -35px;
          border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
                   "
          :show-header="false"
          id="gb-1"
          class="table"
          max-height="250"
      >
        <el-table-column label="头像" prop="avatar">
          <template #default="scope">
            <div style="
        display: flex;
        margin-left: 20px;
        flex-direction: column;
        width: auto;
        height: auto;

      ">
              <div style="display: flex;width: auto;height: auto;" >
                <el-image :src="scope.row.avatar_url" fit="fill" :lazy="true" style="
          width: 50px;
          height: 50px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin-top: 10px;
          margin-right:10px;
          margin-bottom: 10px;
" >
                </el-image>
                <div class="name" style="font-size: 18px; width: auto; white-space: normal;display: flex;flex-direction: row">{{ scope.row.user_name }}
                  <br>
                  {{scope.row.create_time}}
                </div>
              </div>
              <div style="margin-left: 20px; width: auto; white-space: normal">{{ scope.row.context }}</div>
            </div>
          </template>
        </el-table-column>
      </el-table>


      <div class="card" :style="{ 'background-image': 'url('+admins[2].image_url+')' }">   <div class="xbox">
        <el-image :src="admins[2].icon_url" fit="fill" :lazy="true" >
        </el-image>
        <div class="name-box">浠戊</div>
        <div class="good-box">
          <svg  width="45" height="45" viewBox="0 0 20 21" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/></svg>
          <p>114514</p>
        </div>
        <div class="good-box">
          <svg width="45" height="45" viewBox="0 0 19 18" fill="none" :style="{cursor:'pointer'}" id="2" @click="tableview(3)"><path clip-rule="evenodd" d="M.733 2.8a2 2 0 012-2h13.2a2 2 0 012 2v6.015a6 6 0 01-6 6H5.534a4 4 0 00-2.189.652L1.507 16.67a.5.5 0 01-.774-.418V2.8z" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><path d="M4.5 5.62h5" stroke="#333" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></svg>
          <p> 114514</p>
        </div>
      </div>
        <div class="xbox">
          <div class="but-box">
            <div class="desi-box">{{admins[2].title}}</div>
            <el-button type="primary"  @click="hideButtons(3)" class="show" color="#ff661a">撅一下</el-button>
          </div>
          <div class="sing-box">{{admins[2].signature}}</div>
        </div>
        <div class="xbox">
          <el-button type="primary" class="hide" id="3-1" style="display: none" color="#ff661a" ><svg width="24" height="26.666666666666664" viewBox="0 0 18 20" fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M15.882 9.285c.266.33.535.69.653.94l.024.052c.263.555.5 1.054.843 2.069.354 1.046.917 2.977 0 3.31-.867.313-2.035-1.59-2.035-1.59 0 1.147-.64 2.645-2.02 3.726.51.042 1.204.215 1.609.473.405.259.405.628.183.866-.221.238-1.562.23-3.057.217-1.283-.012-2.615-.081-3.239-.146-.656.068-2.06.15-3.368.155-1.385.006-2.467 0-2.745-.226-.277-.226-.228-.593.079-.83s.993-.445 1.53-.507c-1.382-1.083-2.02-2.58-2.02-3.729 0 0-.67 1.896-1.687 1.59-.754-.227-.57-2.271-.275-3.31.24-.839.43-1.27.647-1.758l.158-.362c.107-.25.32-.556.546-.844 1.206.828 2.701 1.109 2.701 1.109l-.15 1.149a.904.904 0 001.787.264l.19-1.15c5.005.475 8.26-.769 9.646-1.468zm-.692-1.14c.003-.115.005-.237.005-.325 0-3.601-1.835-7.22-6.35-7.22C4.327.6 2.49 4.219 2.49 7.819c0 .085.002.201.004.312.68.23 6.853 2.204 12.696.014z" fill="#6FADFF"/></svg></el-button>
          <el-button type="primary" class="hide" id="3-2" style="display: none" :icon="Notebook" color="#ff661a"></el-button>
          <el-button type="primary" class="hide" id="3-3" style="display: none" :icon="Comment" color="#ff661a" @click="dialogVisible=true;whoguest=admins[2].title;guestbookfrom.user_id=3"></el-button>
          <el-button type="primary" class="hide" id="3-4" style="display: none" color="#ff661a" @click="likebut(2)"><svg width="24" height="25.2" viewBox="0 0 20 21" fill="none">
            <path fill-rule="evenodd" clip-rule="evenodd" d="M8.95 2.563l-1.726 3.98-1.158.008v13.748h-.648l6.462.047c3.079-.297 5.037-1.813 5.807-4.473.805-3.99 1.317-6.347 1.538-7.075.366-1.207-.032-2.178-1.235-2.178h-4.699l-.03-3.273c0-1.552-.77-2.358-2.207-2.358-.98 0-1.714.574-2.105 1.574zM4.765 20.294V6.56l-2.14.015a1.8 1.8 0 00-1.788 1.8v10.105a1.8 1.8 0 001.787 1.8l2.14.015z" fill="#FF6880"/>
          </svg></el-button>

        </div></div>
      <el-table
          v-if="table3"
          :data="admins[2].gestbooks"
          style="width: 860px; margin: 0 auto; display: flex; flex-direction: row;margin-left: -35px;
          border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
                   "
          :show-header="false"
          id="gb-1"
          class="table"
          max-height="250"
      >
        <el-table-column label="头像" prop="avatar">
          <template #default="scope">
            <div style="
        display: flex;
        margin-left: 20px;
        flex-direction: column;
        width: auto;
        height: auto;

      ">
              <div style="display: flex;width: auto;height: auto;" >
                <el-image :src="scope.row.avatar_url" fit="fill" :lazy="true" style="
          width: 50px;
          height: 50px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin-top: 10px;
          margin-right:10px;
          margin-bottom: 10px;
" >
                </el-image>
                <div class="name" style="font-size: 18px; width: auto; white-space: normal;display: flex;flex-direction: row">{{ scope.row.user_name }}
                  <br>
                  {{scope.row.create_time}}
                </div>
              </div>
              <div style="margin-left: 20px; width: auto; white-space: normal">{{ scope.row.context }}</div>
            </div>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </template>

  <style scoped lang="scss">
.header{

  height: 800px;
  width: 100%;

}
.header-text{
  position: absolute;
  top: 40%;
  left: 50%;
  transform: translate(-50%,-50%);
}
.heading-primary{
  color:#ffffff;
  text-transform: uppercase;
  font-family: Monocraft;
  backface-visibility: hidden;//解决动画异常摇晃问题
}
.heading-primary-main{
  display: block;
  font-size: 80px;
  font-weight: 400;
  letter-spacing: 25px;
  margin-left: 22.5%;
  animation-name: moveInleft;
  animation-duration: 1s;
  animation-timing-function: ease-out;
  //animation-iteration-count: 3;播放几次
  //animation-delay: 1s;延迟
}
.header-primary-sub{
  display: block;
  font-size: 60px;
  font-weight: 400;
  letter-spacing: 15px;
  animation-name: moveInright;
  animation-duration: 1s;
  animation-timing-function: ease-out;
}
@keyframes moveInleft {
  0%{
    opacity: 0;
    transform: translateX(-100px);
  }
  80%{
    transform: translateX(15px);
  }
  100%{
    opacity:1;
    transform: translate(0);
  }
}
@keyframes moveInright {
  0%{
    opacity: 0;
    transform: translateX(100px);
  }
  80%{
    transform: translateX(-15px);
  }
  100%{
    opacity:1;
    transform: translate(0);
  }
}
.header-arrow {
  position: absolute;
  height: 100px;
  width: 100px;
  top: 80%;
  left: 45%;
  clip-path: polygon(0 50%, 35% 50%, 35% 0, 66% 0, 66% 50%, 100% 51%, 50% 100%);
  background-color: white;
  animation-name: arrowblink;
  animation-duration: 3s;
  animation-timing-function: ease-out;
  animation-iteration-count: infinite; /* Set to 'infinite' for continuous loop */
}

@keyframes arrowblink {
  0%, 100% {
    opacity: 0;
    transform: translateY(-15px);
  }

  50% {
    opacity: 1;
    transform: translateY(0);
  }
}


  .overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5); /* 半透明黑色背景 */
    z-index: 1;
    opacity: 0; /* 默认情况下不可见 */
    pointer-events: none; /* 不可交互 */
    transition: opacity 0.3s;
  }
  .overlay.active {
    opacity: 1; /* 遮罩层激活时可见 */
    pointer-events: auto; /* 可交互 */
  }
  @font-face {
    font-family: ggbond;
    src: url('MFBoHeHaiYan_Noncommercial-Regular.otf')

  }
@font-face {
  font-family: Monocraft;
  src: url("Monocraft.otf");
}
    .card {
      display: flex;
      flex-direction: row;
      align-items: center;
      top: 100px;
      width: 860px;
      height: auto;
      border: 1px solid white;
      box-shadow : 6px 6px 3px rgba(0, 0, 4, 0.3);
      border-radius: 10px;
      min-height: 500px;
      margin-bottom: 150px;
      bottom: 100px;
      position: relative;
     // background-color:transparent;
      background-size: cover;
      z-index: 0;
      margin-left: -50px;
      .good-box{
        width: 150px;
        margin-top: 30px;
        margin-left: 20px;
        align-items: flex-start;
        display: flex;
        flex-direction: row;
        font-size: 48px;
        word-wrap: normal;
        font-family: ggbond;
        color: #7bff91;
      }
      .xbox {
        width: auto;
        height: auto;
        flex-wrap: nowrap;
        .el-image {
          width: 130px;
          height: 130px;
          border-radius: 50%;
          box-shadow: 0 0 4px 2px;
          overflow: hidden;
          margin: 50px 30px;
          margin-top: 0px;
          flex-direction: row;

        }
        .name-box {
          font-family: ggbond;
          word-wrap: break-word;//文本自动换行
          width: 180px;
          text-align: center;
          line-height: 60px;
          font-size: 55px;
          font-weight: bold;
          color: #2fd0be;
        }
        .but-box{
          width: auto;
          height: auto;
          display: flex;
          flex-direction: row;
        }
          .el-button{
            margin-top: 50px;
            margin-left: 150px;
            width: 135px;
            height: 55px;
            border-radius: 20px;
            font-size: 30px;
            font-family: ggbond;
            box-shadow:
                inset 0 -3em 3em rgba(0, 0, 0, 0.1),
                0 0 0 2px rgb(255, 255, 255),
                0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
            transition: all 0.2s;
          }
        @keyframes shake {
          0% {
            transform: translateX(0);
          }
          25% {
            transform: translateX(-5px);
          }
          50% {
            transform: translateX(5px);
          }
          75% {
            transform: translateX(-5px);
          }
          100% {
            transform: translateX(0);
          }
        }
        .el-button.hide:hover{
          transform: scale(1.1); /* 设置按钮获得焦点时的缩放效果，可以根据需要调整 */
          transition: transform 0.2s; /* 添加过渡效果以使动画更平滑 */
        }
        .el-button.show:hover{

          animation: shake 0.5s;
          background-color: #ffd04b;
          //box-shadow: 0 10px 20px ;
        }
        .el-button.show:active{
            transform: translateY(-2px);
        }
        .el-button.hide{
          width: 65px;
          height: 65px;
          border-radius: 25px;
          margin-left: 100px;

        }
        .desi-box{
          background-color: white;
          height: 65px;
          width:300px ;
          margin-top: 50px;
          border-radius: 20px;
          background-color: rgb(153, 204, 255,0.6);
          box-shadow:
              inset 0 -3em 3em rgba(0, 0, 0, 0.1),
              0 0 0 2px rgb(255, 255, 255),
              0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
          text-align: center;
          font-size: 50px;
          font-family: ggbond;
        }

        .sing-box {
          font-family: ggbond;
          margin-top: 50px ;
          width: 600px;
          height: 250px;
          border-radius: 20px;
          background-color: rgba(255,255,255,0.5);
          box-shadow:
              inset 0 -3em 3em rgba(0, 0, 0, 0.1),
              0 0 0 2px rgb(255, 255, 255),
              0.3em 0.3em 1em rgba(0, 0, 0, 0.3);
          font-size: 40px;
        }

        .el-image img {
          width: 100%;
          height: 100%;
          object-fit: cover;
        }
      }
    }
  </style>