// color from dark to light

export const POPULAR = [
  '#ffbe0b',
  '#fb5607',
  '#ff006e',
  '#8338ec',
  '#3a86ff',
  '#06d6a0',
  '#118ab2',
  '#00b4d8',
  '#90e0ef',
  '#caf0f8',
  '#b7e4c7',
  '#c77dff',
  '#f08080',
  '#a1c181',
];

const MAIN_BLUE = ['#7400b8', '#6930c3', '#5e60ce', '#5390d9', '#4ea8de', '#48bfe3', '#56cfe1', '#64dfdf', '#72efdd', '#80ffdb'];

const MAIN_GREEN = ['#10451d', '#155d27', '#1a7431', '#208b3a', '#25a244', '#2dc653', '#4ad66d', '#6ede8a', '#92e6a7', '#b7efc5'];

const MAIN_PURPLE = ['#310055', '#3c0663', '#4a0a77', '#5a108f', '#6818a5', '#8b2fc9', '#ab51e3', '#bd68ee', '#d283ff', '#dc97ff'];

const MAIN_RED = ['#590d22', '#800f2f', '#a4133c', '#c9184a', '#ff4d6d', '#ff758f', '#ff8fa3', '#ffb3c1', '#ffccd5', '#fff0f3'];

const MAIN_BROWSER = ['#583101', '#603808', '#6f4518', '#8b5e34', '#a47148', '#bc8a5f', '#d4a276', '#e7bc91', '#f3d5b5', '#ffedd8'];

const MAIN_YELLOW = ['#ffda0a', '#ffdd1f', '#ffe433', '#ffe747', '#ffec5c', '#ffee70', '#fff185', '#fff599', '#fff8a5', '#ffffb7'];

export const COLOR_PALETTES_MAP = {
  'popular': POPULAR,
  'blue': MAIN_BLUE,
  'green': MAIN_GREEN,
  'purple': MAIN_PURPLE,
  'red': MAIN_RED,
  'browser': MAIN_BROWSER,
  'yellow': MAIN_YELLOW,
};

export const COLOR_PALETTES = [...MAIN_BLUE, ...MAIN_GREEN, ...MAIN_PURPLE, ...MAIN_RED, ...MAIN_BROWSER, ...MAIN_YELLOW];
