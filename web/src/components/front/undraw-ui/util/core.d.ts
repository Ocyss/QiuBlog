/**
 * 防抖
 * 当触发一个函数时，并不会立即执行这个函数，而是会延迟（通过定时器来延迟函数的执行）
 * 如果在延迟时间内，有重新触发函数，那么取消上一次的函数执行（取消定时器）；
 * 如果在延迟时间内，没有重新触发函数，那么这个函数就正常执行（执行传入的函数）
 * debounce(() => console.log('fn 防抖执行了'), 1000)
 * @param fn
 * @param delay
 * @returns
 */
declare const debounce: (fn: (...args: any) => void, delay?: number) => (...args: any) => void;
/**
 * 节流
 * 指定时间持续触发某个事件时，该事件只会执行首次触发事件，也就是说指定时间内只会触发一次
 * @param fn
 * @param duration
 * @returns
 */
declare const throttle: (fn: (...args: any) => void, duration?: number) => (...args: any) => void;
export { debounce, throttle };
