import dividerVue from './divider.vue';
export declare type DividerInstance = InstanceType<typeof dividerVue>;
export declare const UDivider: import('../../util').SFCWithInstall<import("vue").DefineComponent<{
    borderStyle: {
        type: StringConstructor;
        required: false;
        default: string;
    };
    vertical: {
        type: BooleanConstructor;
        required: false;
    };
    position: {
        type: StringConstructor;
        required: false;
        default: string;
    };
}, {
    props: {
        borderStyle: string;
        vertical?: boolean | undefined;
        position: string;
    };
    position: import("vue").Ref<any>;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, {}, string, import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    borderStyle: {
        type: StringConstructor;
        required: false;
        default: string;
    };
    vertical: {
        type: BooleanConstructor;
        required: false;
    };
    position: {
        type: StringConstructor;
        required: false;
        default: string;
    };
}>>, {
    borderStyle: string;
    vertical: boolean;
    position: string;
}>> & Record<string, any>;
export default UDivider;
