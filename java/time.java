import java.util.*;
import java.text.*;
public class time{
	public static void main(String[] args){
		System.out.println(Time("+0d"));
		//System.out.println(getStringDateShort());
	}

	public static String getStringDateShort(Date d) {
	   SimpleDateFormat formatter = new SimpleDateFormat("yyyy-MM-dd");
	   String dateString = formatter.format(d);
	   return dateString;
	}

	//操作时间
	public static Date Time(String s){
		if(!s.matches("^[+-][0-9]+[yMdHms]$")){return new Date();}
		int a=Integer.parseInt(s.split("[yMdHms+-]")[1]);
		Date d=new Date();
		Calendar r=Calendar.getInstance();
		r.set(d.getYear()+1900,d.getMonth(),d.getDate(),d.getHours(),d.getMinutes(),d.getSeconds());
		if(s.matches("[+][0-9]+[y]")){//+年份
			r.add(Calendar.YEAR,a);
		}
		else if(s.matches("[-][0-9]+[y]")){//-年份
			r.add(Calendar.YEAR,-a);
		}
		else if(s.matches("[+][0-9]+[M]")){//+月份
			r.add(Calendar.MONTH,a);
		}
		else if(s.matches("[-][0-9]+[M]")){//-月份
			r.add(Calendar.MONTH,-a);
		}
		else if(s.matches("[+][0-9]+[d]")){//+天数
			r.add(Calendar.DATE,a);
		}
		else if(s.matches("[-][0-9]+[d]")){//-天数
			r.add(Calendar.DATE,-a);
		}
		else if(s.matches("[+][0-9]+[H]")){//+小时
			r.add(Calendar.HOUR,a);
		}
		else if(s.matches("[-][0-9]+[H]")){//-小时
			r.add(Calendar.HOUR,-a);
		}
		else if(s.matches("[+][0-9]+[m]")){//+分钟
			r.add(Calendar.MINUTE,a);
		}
		else if(s.matches("[-][0-9]+[m]")){//-分钟
			r.add(Calendar.MINUTE,-a);
		}
		else if(s.matches("[+][0-9]+[s]")){//+秒数
			r.add(Calendar.SECOND,a);
		}
		else if(s.matches("[-][0-9]+[s]")){//-秒数
			r.add(Calendar.SECOND,-a);
		}
		return r.getTime();
	}
}
