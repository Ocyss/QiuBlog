declare const _sfc_main: import("vue").DefineComponent<{
    name: {
        type: StringConstructor;
        required: false;
    };
    size: {
        type: (NumberConstructor | StringConstructor)[];
        required: false;
    };
    color: {
        type: StringConstructor;
        required: false;
    };
}, {
    props: {
        name?: string | undefined;
        size?: string | number | undefined;
        color?: string | undefined;
    };
    name: import("vue").ComputedRef<string>;
    style: import("vue").ComputedRef<{
        fontSize: string | number | undefined;
        color: string | undefined;
    }>;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    name: {
        type: StringConstructor;
        required: false;
    };
    size: {
        type: (NumberConstructor | StringConstructor)[];
        required: false;
    };
    color: {
        type: StringConstructor;
        required: false;
    };
}>>, {}>;
export default _sfc_main;
