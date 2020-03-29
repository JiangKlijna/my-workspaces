
// @author jiangklijna

class BaseFamily {
	constructor() {
		this.bn = 0 // The number of boys
		this.gn = 0; // The number of girls
    }
	
	isBoy() {
		return Math.random() < 0.5;
	}
}

// 策略一: 生到男孩就不生了,生到女孩,就继续生,直到生到男孩为止
class Family1 extends BaseFamily {
	born() {
		if (super.isBoy()) {
			this.bn++;
			return 
		} else {
			this.gn++;
			this.born();
		}
	}
}

// 策略二: 一定要生两个女儿,男孩无所谓
class Family2 extends BaseFamily {
	born() {
		super.isBoy() ? this.bn++ : this.gn++;
		return (this.gn === 2) ? undefined : this.born();
	}
}

// 策略三: 一个家庭只能要1个孩子,男女无所谓
class Family3 extends BaseFamily {
	born() {
		super.isBoy() ? this.bn++ : this.gn++;
	}
}

// 策略四: 一个家庭有3次生育机会,但每次有四分之一的概率不继续生.
class Family4 extends BaseFamily {
	born() {
		if (Math.random() >= 0.25) super.isBoy() ? this.bn++ : this.gn++;
		if (Math.random() >= 0.25) super.isBoy() ? this.bn++ : this.gn++;
		if (Math.random() >= 0.25) super.isBoy() ? this.bn++ : this.gn++;
	}
}

// 统计一共有Count个Family时,有多少男女
function statistics(Family, Count) {
	let boys = 0, girls = 0;
	Array.from({length: Count}, () => new Family()).forEach((f) => {
		f.born();
		boys += f.bn;
		girls += f.gn;
	});
	return [boys, girls];
}

// main
if (require.main === module) {
	const Family_Number = 1000000;
	const print_statistics = (Family, i) => {
		let [boys, girls] = statistics(Family, Family_Number);
		console.log(`策略${i+1}\t男孩:${boys}, 女孩:${girls}, 男女比例:${(boys/girls).toFixed(4)}`);
	}
	[Family1, Family2, Family3, Family4].forEach(print_statistics);
}
