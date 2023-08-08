import { css } from '@emotion/core';
import React from 'react';
import { Portal } from '../Portal/Portal';

interface IProps {
  isOpen: boolean;
  children: React.ReactElement;
}

export const ModalDialog: React.FC<IProps> = (props) => {
  const { children, isOpen } = props;

  // オープン状態になるたびに新しくPortalを生成することで、
  // 時間的に後に開いたモーダルを<body>直下レベルの「より後方」に描写させる。
  // これにより、z-indexを意識することなく、重なり順を正しく制御することができる。
  if (!isOpen) return null;

  const styles = {
    container: css`
      align-items: center;
      display: flex;
      height: 100vh;
      justify-content: center;
      left: 0;
      position: fixed;
      top: 0;
      width: 100vw;
    `,
    backdrop: css`
      background-color: #000;
      height: 100%;
      opacity: 0.5;
      position: absolute;
      width: 100%;
    `,
    modalContent: css`
      background: white;
      border-radius: 5px;
      max-width: 600px;
      padding: 2rem;
      position: absolute;
      width: 50%;
    `,
  };

  return (
    <Portal>
      <div css={styles.container}>
        <div css={styles.backdrop} />
        <div css={styles.modalContent}>{children}</div>
      </div>
    </Portal>
  );
};
