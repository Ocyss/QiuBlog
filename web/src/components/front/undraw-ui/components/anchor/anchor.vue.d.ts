declare const _sfc_main: import("vue").DefineComponent<{
    container: {
        type: StringConstructor;
        required: true;
    };
    target: {
        type: StringConstructor;
        required: false;
    };
    targetOffset: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
}, {
    props: {
        container: string;
        target?: string | undefined;
        targetOffset: number;
    };
    active: import("vue").Ref<number>;
    navs: import("vue").Ref<{
        [x: number]: HTMLDivElement;
        item: (index: number) => HTMLDivElement;
        forEach: (callbackfn: (value: HTMLDivElement, key: number, parent: NodeListOf<HTMLDivElement>) => void, thisArg?: any) => void;
        readonly length: number;
    }>;
    target: import("vue").Ref<Window | Element>;
    d1: (val: string) => "d2" | "d3" | "d4";
    onScroll: () => void;
    scrollTo: (k: number) => void;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    container: {
        type: StringConstructor;
        required: true;
    };
    target: {
        type: StringConstructor;
        required: false;
    };
    targetOffset: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
}>>, {
    targetOffset: number;
}>;
export default _sfc_main;
