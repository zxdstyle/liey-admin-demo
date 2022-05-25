interface Window {
  $loadingBar?: import('naive-ui').LoadingBarProviderInst;
  $dialog?: import('naive-ui').DialogProviderInst;
  $message?: import('naive-ui').MessageProviderInst;
  $notification?: import('naive-ui').NotificationProviderInst;
}

/** 通用类型 */
declare namespace Common {
  /**
   * 策略模式
   * [状态, 为true时执行的回调函数]
   */
  type StrategyAction = [boolean, () => void];
}

/** 构建时间 */
declare const PROJECT_BUILD_TIME: string;

declare type Nullable<T> = T | null;
declare type Recordable<T = any> = Record<string, T>;
declare type EmitType = (event: 'register', ...args: any[]) => void;
declare interface Fn<T = any, R = T> {
  (...arg: T[]): R;
}

declare type TargetContext = '_self' | '_blank';

declare interface ComponentElRef<T extends HTMLElement = HTMLDivElement> {
  $el: T;
}

declare type ComponentRef<T extends HTMLElement = HTMLDivElement> = ComponentElRef<T> | null;

declare type ElRef<T extends HTMLElement = HTMLDivElement> = Nullable<T>;
