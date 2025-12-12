// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/*
一个mapping来存储候选人的得票数
一个vote函数，允许用户投票给某个候选人
一个getVotes函数，返回某个候选人的得票数
一个resetVotes函数，重置所有候选人的得票数
*/
contract Practice {
    mapping(uint => string) mp1;
    mapping(uint => string) mp2;
    mapping(string => uint) mp3;
    mapping(string => uint) mp4;
    
    constructor() {
        mp1[4] = "IV";
        mp1[9] = "IX";
        mp1[40] = "XL";
        mp1[90] = "XC";
        mp1[400] = "CD";
        mp1[900] = "CM";

        mp2[1] = "I";
        mp2[5] = "V";
        mp2[10] = "X";
        mp2[50] = "L";
        mp2[100] = "C";
        mp2[500] = "D";
        mp2[1000] = "M";

        mp3["IV"] = 4;
        mp3["IX"] = 9;
        mp3["XL"] = 40;
        mp3["XC"] = 90;
        mp3["CD"] = 400;
        mp3["CM"] = 900;

        mp4["I"] = 1;
        mp4["V"] = 5;
        mp4["X"] = 10;
        mp4["L"] = 50;
        mp4["C"] = 100;
        mp4["D"] = 500;
        mp4["M"] = 1000;
    }

    //2.反转字符串
    function revertString(string memory ss) public pure returns (string memory) {
        bytes memory str = bytes(ss);
        uint len = str.length;
        bytes memory result = new bytes(len);
        for (uint i = 0; i < len; i++) 
        {
            result[i] = str[len - i - 1];
        }
        return string(result);
    }



    // //3.用 solidity 实现整数转罗马数字
    // //题目描述在 https://leetcode.cn/problems/roman-to-integer/description/3.
    function numberToRoma(uint num) public view returns (string memory) {
        uint base = 1;
        string memory result;
        while(num > 0){
            uint r = num % 10;
            uint tmp = r * base;
            if(bytes(mp1[tmp]).length != 0) {
                result = string(abi.encodePacked(mp1[tmp], result));
            } else if(r == 5) {
                result = string(abi.encodePacked(mp2[tmp], result));
            } else if(r > 5) {
                result = string(abi.encodePacked(mp2[5 * base], numberToRomaHelper(r-5, base), result));
            } else {
                result = string(abi.encodePacked(numberToRomaHelper(r, base), result));
            }

            base *= 10;
            num /= 10;
        }
        return result;
    }

    function numberToRomaHelper(uint cnt, uint base) private view returns (string memory) {
        string memory res;
        for (uint i = 0; i < cnt; i++) 
        {
            res = string(abi.encodePacked(res, mp2[base]));
        }
        return res;
    }
    // //4.用 solidity 实现罗马数字转数整数
    // //题目描述在 https://leetcode.cn/problems/integer-to-roman/description/
    function romaToNumer(string memory s) public view returns (uint) {
        uint result;
        uint i;
        bytes memory sb = bytes(s);
        bytes memory tmp2 = new bytes(1);
        for (i = 0; i < sb.length - 1; ){
            bytes1 c1 = sb[i];
            bytes1 c2 = sb[i + 1];

            bytes memory tmp = new bytes(2);
            tmp[0] = c1;
            tmp[1] = c2;

            tmp2[0] = c1;

            if(mp3[string(tmp)] != 0) {
                i += 2;
                result += mp3[string(tmp)];
            } else {
                result += mp4[string(tmp2)];
                i++;
            }
        }

        if(i == sb.length - 1) {
            result += mp4[string(tmp2)];
        }

	    return result;
    }

    //5.合并两个有序数组 (Merge Sorted Array)
    //题目描述：将两个有序数组合并为一个有序数组。
    function mergeSortedArray(uint[] calldata arr1, uint[] calldata arr2) public pure returns (uint[] memory) {
        uint len1 = arr1.length;
        uint len2 = arr2.length;
        uint[] memory result = new uint[](len1 + len2);
        uint i;
        uint j;
        uint k;
        while(i < len1 && j < len2) {
            if(arr1[i] > arr2[j]) {
                result[k] = arr2[j];
                j++;
            } else {
                result[k] = arr1[i];
                i++;
            }
            k++;
        }
        if(i == len1) {
            for (; j < len2; j++) 
            {
                result[k] = arr2[j];
                k++;
            }
        } else {
            for (; i < len1; i++) 
            {
                result[k] = arr1[i];
                k++;
            }
        }
        return result;
    
    }

    //6.二分查找 (Binary Search)
    //题目描述：在一个有序数组中查找目标值。
    function binarySearch(uint[] calldata arr, uint target) public pure returns (uint,bool) {
        uint len = arr.length;
        uint mid;
        uint l = 0;
        uint r = len - 1;
        while(l <= r) {
            mid = (l + r) / 2;
            if(arr[mid] == target) {
                return (mid,true);
            } else if(arr[mid] > target) {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }
        return (0,false);
    }
    
}