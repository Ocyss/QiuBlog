import { TagApi } from '.';
declare const _sfc_main: import("vue").DefineComponent<{
    dropdown: {
        type: ObjectConstructor;
        required: true;
    };
}, {
    props: {
        dropdown: {
            x: number;
            y: number;
        };
    };
    state: {
        tag: {
            path: string;
            name?: string | undefined;
            meta: {
                title: string;
                isAffix: boolean;
                isKeepAlive?: boolean | undefined;
            };
        };
        isShow: boolean;
        dropdownList: {
            title: string;
            show: boolean;
            icon: string;
        }[];
    };
    emit: (e: 'submit', val: number, tag: TagApi) => void;
    openContextmenu: (v: TagApi) => void;
    closeContextmenu: () => void;
    isShow: import("vue").Ref<boolean>;
    dropdownList: import("vue").Ref<{
        title: string;
        show: boolean;
        icon: string;
    }[]>;
    tag: import("vue").Ref<{
        path: string;
        name?: string | undefined;
        meta: {
            title: string;
            isAffix: boolean;
            isKeepAlive?: boolean | undefined;
        };
    }>;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, "submit"[], "submit", import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    dropdown: {
        type: ObjectConstructor;
        required: true;
    };
}>> & {
    onSubmit?: ((...args: any[]) => any) | undefined;
}, {}>;
export default _sfc_main;
