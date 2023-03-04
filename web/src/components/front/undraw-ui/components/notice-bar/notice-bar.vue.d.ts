/// <reference types="node" />
declare const _sfc_main: import("vue").DefineComponent<{
    data: {
        type: null;
        required: true;
    };
    size: {
        type: (NumberConstructor | StringConstructor)[];
        required: false;
        default: number;
    };
    vertical: {
        type: BooleanConstructor;
        required: false;
    };
    height: {
        type: (NumberConstructor | StringConstructor)[];
        required: false;
        default: number;
    };
    delay: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    spped: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    suffixIcon: {
        type: StringConstructor;
        required: false;
    };
    prefixIcon: {
        type: StringConstructor;
        required: false;
    };
    color: {
        type: StringConstructor;
        required: false;
        default: string;
    };
    background: {
        type: StringConstructor;
        required: false;
        default: string;
    };
}, {
    props: {
        data: any;
        size: number | string;
        vertical?: boolean | undefined;
        height: number | string;
        delay: number;
        spped: number;
        suffixIcon?: string | undefined;
        prefixIcon?: string | undefined;
        color: string;
        background: string;
    };
    state: {
        boxWidth: number;
        textWidth: number;
        oneTime: number;
        twoTime: number;
        order: number;
    };
    boxRef: import("vue").Ref<HTMLDivElement>;
    textRef: import("vue").Ref<HTMLDivElement>;
    interval: import("vue").ComputedRef<number>;
    initAnimation: () => void;
    computeAnimationTime: () => void;
    changeAnimation: () => void;
    listenerAnimationend: () => void;
    ElCarousel: import("element-plus/es/utils").SFCWithInstall<import("vue").DefineComponent<{
        readonly initialIndex: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, 0, boolean>;
        readonly height: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
        readonly trigger: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "hover" | "click", unknown, "hover", boolean>;
        readonly autoplay: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        readonly interval: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, 3000, boolean>;
        readonly indicatorPosition: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "" | "none" | "outside", unknown, "", boolean>;
        readonly indicator: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        readonly arrow: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "hover" | "always" | "never", unknown, "hover", boolean>;
        readonly type: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "" | "card", unknown, "", boolean>;
        readonly loop: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        readonly direction: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "vertical" | "horizontal", unknown, "horizontal", boolean>;
        readonly pauseOnHover: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
    }, {
        props: Readonly<import("@vue/shared").LooseRequired<Readonly<import("vue").ExtractPropTypes<{
            readonly initialIndex: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, 0, boolean>;
            readonly height: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
            readonly trigger: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "hover" | "click", unknown, "hover", boolean>;
            readonly autoplay: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
            readonly interval: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, 3000, boolean>;
            readonly indicatorPosition: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "" | "none" | "outside", unknown, "", boolean>;
            readonly indicator: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
            readonly arrow: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "hover" | "always" | "never", unknown, "hover", boolean>;
            readonly type: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "" | "card", unknown, "", boolean>;
            readonly loop: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
            readonly direction: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "vertical" | "horizontal", unknown, "horizontal", boolean>;
            readonly pauseOnHover: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        }>> & {
            onChange?: ((current: number, prev: number) => any) | undefined;
        }>>;
        emit: (event: "change", current: number, prev: number) => void;
        ns: {
            namespace: import("vue").Ref<string>;
            b: (blockSuffix?: string | undefined) => string;
            e: (element?: string | undefined) => string;
            m: (modifier?: string | undefined) => string;
            be: (blockSuffix?: string | undefined, element?: string | undefined) => string;
            em: (element?: string | undefined, modifier?: string | undefined) => string;
            bm: (blockSuffix?: string | undefined, modifier?: string | undefined) => string;
            bem: (blockSuffix?: string | undefined, element?: string | undefined, modifier?: string | undefined) => string;
            is: {
                (name: string, state: boolean | undefined): string;
                (name: string): string;
            };
            cssVar: (object: Record<string, string>) => Record<string, string>;
            cssVarName: (name: string) => string;
            cssVarBlock: (object: Record<string, string>) => Record<string, string>;
            cssVarBlockName: (name: string) => string;
        };
        COMPONENT_NAME: string;
        THROTTLE_TIME: number;
        activeIndex: import("vue").Ref<number>;
        timer: import("vue").Ref<{
            hasRef: () => boolean;
            refresh: () => NodeJS.Timer;
            ref: () => NodeJS.Timer;
            unref: () => NodeJS.Timer;
            [Symbol.toPrimitive]: () => number;
        } | null>;
        hover: import("vue").Ref<boolean>;
        root: import("vue").Ref<HTMLDivElement | undefined>;
        items: import("vue").Ref<{
            props: {
                readonly name: string;
                readonly label: import("element-plus/es/utils").EpPropMergeType<readonly [StringConstructor, NumberConstructor], unknown, unknown>;
            };
            states: {
                hover: boolean;
                translate: number;
                scale: number;
                active: boolean;
                ready: boolean;
                inStage: boolean;
                animating: boolean;
            };
            uid: number | undefined;
            translateItem: (index: number, activeIndex: number, oldIndex?: number | undefined) => void;
        }[]>;
        arrowDisplay: import("vue").ComputedRef<boolean>;
        hasLabel: import("vue").ComputedRef<boolean>;
        carouselClasses: import("vue").ComputedRef<string[]>;
        indicatorsClasses: import("vue").ComputedRef<string[]>;
        isCardType: import("vue").ComputedRef<boolean>;
        isVertical: import("vue").ComputedRef<boolean>;
        throttledArrowClick: import("lodash").DebouncedFunc<(index: number) => void>;
        throttledIndicatorHover: import("lodash").DebouncedFunc<(index: number) => void>;
        pauseTimer: () => void;
        startTimer: () => void;
        playSlides: () => void;
        setActiveItem: (index: string | number) => void;
        resetItemPosition: (oldIndex?: number | undefined) => void;
        addItem: (item: import("element-plus").CarouselItemContext) => void;
        removeItem: (uid?: number | undefined) => void;
        itemInStage: (item: import("element-plus").CarouselItemContext, index: number) => false | "right" | "left";
        handleMouseEnter: () => void;
        handleMouseLeave: () => void;
        handleButtonEnter: (arrow: "right" | "left") => void;
        handleButtonLeave: () => void;
        handleIndicatorClick: (index: number) => void;
        handleIndicatorHover: (index: number) => void;
        prev: () => void;
        next: () => void;
        resetTimer: () => void;
        resizeObserver: import("vue").ShallowRef<{
            isSupported: import("vue").Ref<boolean>;
            stop: () => void;
        } | undefined>;
        ElIcon: import("element-plus/es/utils").SFCWithInstall<import("vue").DefineComponent<{
            readonly size: {
                readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => (string | number) & {}) | (() => string | number) | ((new (...args: any[]) => (string | number) & {}) | (() => string | number))[], unknown, unknown>>;
                readonly required: false;
                readonly validator: ((val: unknown) => boolean) | undefined;
                __epPropKey: true;
            };
            readonly color: {
                readonly type: import("vue").PropType<string>;
                readonly required: false;
                readonly validator: ((val: unknown) => boolean) | undefined;
                __epPropKey: true;
            };
        }, {
            props: Readonly<import("@vue/shared").LooseRequired<Readonly<import("vue").ExtractPropTypes<{
                readonly size: {
                    readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => (string | number) & {}) | (() => string | number) | ((new (...args: any[]) => (string | number) & {}) | (() => string | number))[], unknown, unknown>>;
                    readonly required: false;
                    readonly validator: ((val: unknown) => boolean) | undefined;
                    __epPropKey: true;
                };
                readonly color: {
                    readonly type: import("vue").PropType<string>;
                    readonly required: false;
                    readonly validator: ((val: unknown) => boolean) | undefined;
                    __epPropKey: true;
                };
            }>> & {
                [x: string & `on${string}`]: ((...args: any[]) => any) | ((...args: unknown[]) => any) | undefined;
            }>>;
            ns: {
                namespace: import("vue").Ref<string>;
                b: (blockSuffix?: string | undefined) => string;
                e: (element?: string | undefined) => string;
                m: (modifier?: string | undefined) => string;
                be: (blockSuffix?: string | undefined, element?: string | undefined) => string;
                em: (element?: string | undefined, modifier?: string | undefined) => string;
                bm: (blockSuffix?: string | undefined, modifier?: string | undefined) => string;
                bem: (blockSuffix?: string | undefined, element?: string | undefined, modifier?: string | undefined) => string;
                is: {
                    (name: string, state: boolean | undefined): string;
                    (name: string): string;
                };
                cssVar: (object: Record<string, string>) => Record<string, string>;
                cssVarName: (name: string) => string;
                cssVarBlock: (object: Record<string, string>) => Record<string, string>;
                cssVarBlockName: (name: string) => string;
            };
            style: import("vue").ComputedRef<import("vue").CSSProperties>;
        }, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, Record<string, any>, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
            readonly size: {
                readonly type: import("vue").PropType<import("element-plus/es/utils").EpPropMergeType<(new (...args: any[]) => (string | number) & {}) | (() => string | number) | ((new (...args: any[]) => (string | number) & {}) | (() => string | number))[], unknown, unknown>>;
                readonly required: false;
                readonly validator: ((val: unknown) => boolean) | undefined;
                __epPropKey: true;
            };
            readonly color: {
                readonly type: import("vue").PropType<string>;
                readonly required: false;
                readonly validator: ((val: unknown) => boolean) | undefined;
                __epPropKey: true;
            };
        }>>, {}>> & Record<string, any>;
        ArrowLeft: import("vue").DefineComponent<{}, {}, {}, import("vue").ComputedOptions, import("vue").MethodOptions, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{}>>, {}>;
        ArrowRight: import("vue").DefineComponent<{}, {}, {}, import("vue").ComputedOptions, import("vue").MethodOptions, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{}>>, {}>;
    }, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {
        change: (current: number, prev: number) => boolean;
    }, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
        readonly initialIndex: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, 0, boolean>;
        readonly height: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
        readonly trigger: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "hover" | "click", unknown, "hover", boolean>;
        readonly autoplay: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        readonly interval: import("element-plus/es/utils").EpPropFinalized<NumberConstructor, unknown, unknown, 3000, boolean>;
        readonly indicatorPosition: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "" | "none" | "outside", unknown, "", boolean>;
        readonly indicator: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        readonly arrow: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "hover" | "always" | "never", unknown, "hover", boolean>;
        readonly type: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "" | "card", unknown, "", boolean>;
        readonly loop: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
        readonly direction: import("element-plus/es/utils").EpPropFinalized<StringConstructor, "vertical" | "horizontal", unknown, "horizontal", boolean>;
        readonly pauseOnHover: import("element-plus/es/utils").EpPropFinalized<BooleanConstructor, unknown, unknown, true, boolean>;
    }>> & {
        onChange?: ((current: number, prev: number) => any) | undefined;
    }, {
        readonly type: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "" | "card", unknown>;
        readonly trigger: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "hover" | "click", unknown>;
        readonly height: string;
        readonly direction: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "vertical" | "horizontal", unknown>;
        readonly loop: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
        readonly indicator: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
        readonly initialIndex: number;
        readonly autoplay: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
        readonly interval: number;
        readonly indicatorPosition: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "" | "none" | "outside", unknown>;
        readonly arrow: import("element-plus/es/utils").EpPropMergeType<StringConstructor, "hover" | "always" | "never", unknown>;
        readonly pauseOnHover: import("element-plus/es/utils").EpPropMergeType<BooleanConstructor, unknown, unknown>;
    }>> & {
        CarouselItem: import("vue").DefineComponent<{
            readonly name: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
            readonly label: import("element-plus/es/utils").EpPropFinalized<readonly [StringConstructor, NumberConstructor], unknown, unknown, "", boolean>;
        }, {
            props: Readonly<import("@vue/shared").LooseRequired<Readonly<import("vue").ExtractPropTypes<{
                readonly name: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
                readonly label: import("element-plus/es/utils").EpPropFinalized<readonly [StringConstructor, NumberConstructor], unknown, unknown, "", boolean>;
            }>> & {
                [x: string & `on${string}`]: ((...args: any[]) => any) | ((...args: unknown[]) => any) | undefined;
            }>>;
            ns: {
                namespace: import("vue").Ref<string>;
                b: (blockSuffix?: string | undefined) => string;
                e: (element?: string | undefined) => string;
                m: (modifier?: string | undefined) => string;
                be: (blockSuffix?: string | undefined, element?: string | undefined) => string;
                em: (element?: string | undefined, modifier?: string | undefined) => string;
                bm: (blockSuffix?: string | undefined, modifier?: string | undefined) => string;
                bem: (blockSuffix?: string | undefined, element?: string | undefined, modifier?: string | undefined) => string;
                is: {
                    (name: string, state: boolean | undefined): string;
                    (name: string): string;
                };
                cssVar: (object: Record<string, string>) => Record<string, string>;
                cssVarName: (name: string) => string;
                cssVarBlock: (object: Record<string, string>) => Record<string, string>;
                cssVarBlockName: (name: string) => string;
            };
            COMPONENT_NAME: string;
            carouselContext: import("element-plus").CarouselContext;
            instance: import("vue").ComponentInternalInstance;
            CARD_SCALE: number;
            hover: import("vue").Ref<boolean>;
            translate: import("vue").Ref<number>;
            scale: import("vue").Ref<number>;
            active: import("vue").Ref<boolean>;
            ready: import("vue").Ref<boolean>;
            inStage: import("vue").Ref<boolean>;
            animating: import("vue").Ref<boolean>;
            isCardType: import("vue").Ref<boolean>;
            isVertical: import("vue").Ref<boolean>;
            itemStyle: import("vue").ComputedRef<import("vue").CSSProperties>;
            processIndex: (index: number, activeIndex: number, length: number) => number;
            calcCardTranslate: (index: number, activeIndex: number) => number;
            calcTranslate: (index: number, activeIndex: number, isVertical: boolean) => number;
            translateItem: (index: number, activeIndex: number, oldIndex?: number | undefined) => void;
            handleItemClick: () => void;
        }, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, Record<string, any>, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
            readonly name: import("element-plus/es/utils").EpPropFinalized<StringConstructor, unknown, unknown, "", boolean>;
            readonly label: import("element-plus/es/utils").EpPropFinalized<readonly [StringConstructor, NumberConstructor], unknown, unknown, "", boolean>;
        }>>, {
            readonly name: string;
            readonly label: import("element-plus/es/utils").EpPropMergeType<readonly [StringConstructor, NumberConstructor], unknown, unknown>;
        }>;
    };
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    data: {
        type: null;
        required: true;
    };
    size: {
        type: (NumberConstructor | StringConstructor)[];
        required: false;
        default: number;
    };
    vertical: {
        type: BooleanConstructor;
        required: false;
    };
    height: {
        type: (NumberConstructor | StringConstructor)[];
        required: false;
        default: number;
    };
    delay: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    spped: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    suffixIcon: {
        type: StringConstructor;
        required: false;
    };
    prefixIcon: {
        type: StringConstructor;
        required: false;
    };
    color: {
        type: StringConstructor;
        required: false;
        default: string;
    };
    background: {
        type: StringConstructor;
        required: false;
        default: string;
    };
}>>, {
    size: string | number;
    color: string;
    background: string;
    vertical: boolean;
    height: string | number;
    delay: number;
    spped: number;
}>;
export default _sfc_main;
