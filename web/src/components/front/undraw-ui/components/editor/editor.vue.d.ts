declare const _sfc_main: import("vue").DefineComponent<{
    placeholder: {
        type: StringConstructor;
        required: false;
    };
    modelValue: {
        type: StringConstructor;
        required: true;
    };
    minHeight: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    imgList: {
        type: ArrayConstructor;
        required: false;
    };
}, {
    props: {
        placeholder?: string | undefined;
        modelValue: string;
        minHeight: number;
        imgList?: string[] | undefined;
    };
    range: import("vue").Ref<Range | undefined>;
    editorRef: import("vue").Ref<HTMLDivElement | undefined>;
    text: import("vue").Ref<any>;
    isLocked: import("vue").Ref<boolean>;
    active: import("vue").Ref<boolean>;
    imageRef: import("vue").Ref<HTMLDivElement | undefined>;
    imgList: import("vue").Ref<string[] | undefined> | undefined;
    minHeight: import("vue").ComputedRef<string>;
    padding: import("vue").ComputedRef<"4px 10px" | "8px 12px">;
    emit: {
        (e: 'update:modelValue', val: string): void;
        (e: 'input', event: Event): void;
        (e: 'focus', event: Event): void;
        (e: 'blur', event: Event): void;
        (e: 'submit'): void;
    };
    onFocus: (event: Event) => void;
    onBlur: (event: Event) => void;
    onInput: (event: Event) => void;
    addText: (val: string) => void;
    clear: () => void;
    focus: () => void;
    keyDown: (e: KeyboardEvent) => void;
    removeImg: (val: number) => void;
}, unknown, {}, {}, import("vue").ComponentOptionsMixin, import("vue").ComponentOptionsMixin, ("submit" | "update:modelValue" | "input" | "focus" | "blur")[], "submit" | "update:modelValue" | "input" | "focus" | "blur", import("vue").VNodeProps & import("vue").AllowedComponentProps & import("vue").ComponentCustomProps, Readonly<import("vue").ExtractPropTypes<{
    placeholder: {
        type: StringConstructor;
        required: false;
    };
    modelValue: {
        type: StringConstructor;
        required: true;
    };
    minHeight: {
        type: NumberConstructor;
        required: false;
        default: number;
    };
    imgList: {
        type: ArrayConstructor;
        required: false;
    };
}>> & {
    onSubmit?: ((...args: any[]) => any) | undefined;
    "onUpdate:modelValue"?: ((...args: any[]) => any) | undefined;
    onInput?: ((...args: any[]) => any) | undefined;
    onFocus?: ((...args: any[]) => any) | undefined;
    onBlur?: ((...args: any[]) => any) | undefined;
}, {
    minHeight: number;
}>;
export default _sfc_main;
