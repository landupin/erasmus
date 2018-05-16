import {MDCRipple} from '@material/ripple';
import {MDCToolbar} from '@material/toolbar';
import {MDCTemporaryDrawer} from '@material/drawer';

const ripple = new MDCRipple(document.querySelector('.button'));
const toolbar = new MDCToolbar(document.querySelector('.mdc-toolbar'));

const drawer = new MDCTemporaryDrawer(document.querySelector('.mdc-drawer--temporary'));
document.querySelector('.menu').addEventListener('click', () => drawer.open = true);