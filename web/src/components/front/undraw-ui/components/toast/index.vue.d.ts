export interface ToastApi {
    message?: string;
    duration?: number;
    type?: string;
}
declare const _sfc_main: import("vue").DefineComponent<{
    message: {
        type: StringConstructor;
        required: false;
        default: string;
    };
    duration: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    type: {
        type: StringConstructor;
        required: false;
        default: string;
    };
}, {
    props: {
        message: string;
        duration: number;
        type: string;
    };
    alert: {
        color: string;
        bgColor: string;
        icon: string;
    };
    visible: import("vue").Ref<boolean>;
    iconVue: import("../..").SFCWithInstall<import("vue").DefineComponent<{
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
    }>>, {}>> & Record<string, any>;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    message: {
        type: StringConstructor;
        required: false;
        default: string;
    };
    duration: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    type: {
        type: StringConstructor;
        required: false;
        default: string;
    };
}>>, {
    type: string;
    duration: number;
    message: string;
}>;
export default _sfc_main;
