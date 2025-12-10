// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Rome{

     bytes bts ;
//     字符          数值
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
   
    //给定一个罗马数字，将其转换成整数。
    function getUint(string memory str) public returns(uint256) {
        bytes memory bt = bytes(str);
        bytes memory i1= bytes("I");
        bytes memory v5= bytes("V");
        bytes memory x10 = bytes("X");
        bytes memory l50=bytes("L");
        bytes memory c100=bytes("C");
        bytes memory d500 = bytes("D");
        bytes memory m1000 = bytes("M");
        string memory s;
        uint256 res;
        uint256 len = bt.length;
        for(uint256 i = 0; i < len; i++){
            bytes memory tem = new bytes(1);
            tem[0] = bt[len-1-i];
            //string memory lt = string(tem);
            if(i1[0] == tem[0]){
                if (res ==5 || res ==10){
                    res -=1;
                }else{
                   res+=1; 
                }
                
            }
            if(v5[0] == tem[0]){
                res+=5;
            }
             if(x10[0] == tem[0]){
                if ( res >= 50 || res >= 100){
                    res -=10;
                }else{
                   res+=10; 
                }
            }

             if(l50[0] == tem[0]){
                res+=50;
            }
            if(c100[0] == tem[0]){
                 if ( res >= 500 || res >= 1000){
                    res -=100;
                }else{
                   res+=100; 
                }
                
            }
             if(d500[0] == tem[0]){
                res+=500;
            }
            if(m1000[0] == tem[0]){
                res+=1000;
            }
        }
        return res;
    }
    
    function getRomeNumber(uint256  input) public   returns(string memory ) {

        uint256 u = input / 1000;
        uint256 m = input % 1000;
        uint256 d = m/100;
        uint256 c = m % 100;
        uint256 v= c / 10;
        uint256 g= c % 10;
        bytes storage bt = bts ;

        if (u >= 1){
            for (uint256 i = 0; i< u;i++){
                 bt.push(bytes1("M"));
            }
        }
        if(d == 9){
            bt.push(bytes1("C"));
            bt.push(bytes1("M"));
        }
         if(d >5 && d <9){
            bt.push(bytes1("D"));
            uint tm = d -5;
            for (uint256 i = 0; i< tm;i++){
                 bt.push(bytes1("C"));
            }
        } 
        if (d == 5){
            bt.push(bytes1("D"));
        }
        if (d == 4){
            bt.push(bytes1("C"));
            bt.push(bytes1("D"));
        }
         if(d <=3){
            uint tm = d;
            for (uint256 i = 0; i< tm;i++){
                 bt.push(bytes1("C"));
            }
        } 


        if(v == 9){
            bt.push(bytes1("X"));
            bt.push(bytes1("C"));
        }
       
        if(v >5 && v <9){
            bt.push(bytes1("L"));
            uint tm = v -5;
            for (uint256 i = 0; i< tm;i++){
                 bt.push(bytes1("X"));
            }
        } 
        if (v == 5){
            bt.push(bytes1("L"));
        }
        if (v == 4){
            bt.push(bytes1("X"));
            bt.push(bytes1("L"));
        }
         if(v <=3){
            uint tm = v;
            for (uint256 i = 0; i< tm;i++){
                 bt.push(bytes1("X"));
            }
        } 



        if(g == 9){
            bt.push(bytes1("I"));
            bt.push(bytes1("X"));
        }
        if(g >5 && g <9){
            bt.push(bytes1("V"));
            uint tm = g -5;
            for (uint256 i = 0; i< tm;i++){
                 bt.push(bytes1("I"));
            }
        } 
        if (g == 5){
            bt.push(bytes1("V"));
        }
        if (g == 4){
            bt.push(bytes1("I"));
            bt.push(bytes1("V"));
        }
         if(g <=3){
            uint tm = g;
            for (uint256 i = 0; i< tm;i++){
                 bt.push(bytes1("I"));
            }
        } 


        return string(bt);
    }

    function newOrder(uint256[] memory s1, uint256[] memory s2) public pure returns (uint256[] memory) {
        uint256 len1 = s1.length;
        uint256 len2 = s2.length;
        uint256 len3 = len1 + len2;
        uint256[] memory news = new uint256[](len3); // 初始化结果数组
        uint256 a = 0; // s1 的遍历指针
        uint256 b = 0; // s2 的遍历指针
        uint256 i = 0; // 结果数组的索引指针

        while (a < len1 && b < len2) {
            if (s1[a] < s2[b]) { // 直接用指针访问，无需临时变量 t1/t2
                news[i] = s1[a];
                a++;
            } else {
                news[i] = s2[b];
                b++;
            }
            i++;
        }

        // 场景2：s1 还有剩余元素，全部追加
        while (a < len1) {
            news[i] = s1[a];
            a++;
            i++;
        }

        // 场景3：s2 还有剩余元素，全部追加
        while (b < len2) {
            news[i] = s2[b];
            b++;
            i++;
        }

        return news;
    }  

//末尾更新 mid = left + (right-left)/2 = 2 + (1-2)/2 → 
//注意：right=1 < left=2，(1-2) 是 uint256 下溢，结果为 2^256-1，
//所以 mid = 2 + (2^256-1)/2 ≈ 2^255（极大值）；
    function binarySearch(uint256[] memory s1,uint256 t) external  pure returns (bool){
        uint256 len = s1.length;
        uint256 left =0;
        uint256 right = len -1;
        uint256 mid = left + (right - left)/2;
        //uint256 p ;
        while(left <= right){
             mid = left + (right - left)/2;
            if(s1[mid] == t){
                return true;
            }
            if(s1[mid] > t){
                right = mid - 1;
            }else {
                left = mid + 1;
            }
            
        }
        return false;
    }
} 
