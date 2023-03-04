/**
 * @param data
 * @returns
 *
 * 适用于任何数组和对象的深拷贝
 * const arr = [1, 2, 3, { a: "b" }, [{ c: "d" }]];
 * const cloneArr = cloneDeep(arr);
 *
 * cloneArr[3].a = "u";
 * cloneArr[4][0].c = "uu";
 *
 * console.log(arr); // [1, 2, 3, { a: "b" }, [{ c: "d" }]]
 * console.log(cloneArr); // [1, 2, 3, { a: "u" }, [{ c: "uu" }]]
 */
export declare function cloneDeep(data: any): any;
/**
 * 数组结构转换树型结构
 * @param list
 * @param param: { parentId: string, children: string } //parentId 判断条件属性名， children转换树型结构合并成数组的属性名
 * @returns
 */
export declare function deepTree(list: any[], { parentId, children }: {
    parentId?: string | undefined;
    children?: string | undefined;
}): any[];
/**
 * 树型结构转换为数组结构
 * @param list
 * @param param: { parentId: string, children: string} //parentId 判断条件属性名, 默认值parentId， children转换数据结构拆分对象的属性名 默认值: children
 * @returns
 */
export declare function revDeepTree(list: any[] | undefined, { parentId, children }: {
    parentId?: string | undefined;
    children?: string | undefined;
}): any[];
/**
 *
 * @param arr
 * @param depth
 * @returns
 *
 * 无穷数(Infinity)
 * 将array递归为一维数组。
 * const array = [1,2,[1,2]];
 * console.log(array.flattenDeep(Infinity));  // [1,2,1,2]
 */
export declare const flattenDeep: (arr: any[], depth?: number) => any[];
