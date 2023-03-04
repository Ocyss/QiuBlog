import iconVue from './icon.vue';
export declare type IconInstance = InstanceType<typeof iconVue>;
export declare const UIcon: import('../../util').SFCWithInstall<import("vue").DefineComponent<{
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
export default UIcon;
