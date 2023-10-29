import mapObj from 'map-obj';

export const camelCase = (str: string) => {
  return str.replaceAll(/[._-](\w|$)/g, (_, x) => x.toUpperCase());
};

export const snakeCase = (str: string) => {
  return str
    .replaceAll(/^[^\dA-Za-z]*|[^\dA-Za-z]*$/g, '')
    .replaceAll(/([a-z])([A-Z])/g, (m, a, b) => `${a}_${b.toLowerCase()}`)
    .replaceAll(/[^\dA-Za-z]+|_+/g, '_')
    .toLowerCase();
};

export const kebabCase = (str: string) => {
  return str
    .toString()
    .match(/[A-Z]{2,}(?=[A-Z][a-z]+\d*|\b)|[A-Z]?[a-z]+\d*|[A-Z]|\d+/g)
    .map((x) => x.toLowerCase())
    .join('-');
};

type changeCaseType = 'toCamelCase' | 'toSnakeCase';

const isObject = (v) => typeof v === 'object';

const changer = (str: string, type: changeCaseType) => {
  if (type == 'toCamelCase') {
    return camelCase(str);
  }
  if (type == 'toSnakeCase') {
    return snakeCase(str);
  }
  return null;
};

function objectValues(obj) {
  return Object.keys(obj).map((key) => {
    return obj[key];
  });
}

const changeStringCaseRecursive = (obj: object, type: changeCaseType, isRecursiveChild = false) => {
  const transform =
    obj == null || (Array.isArray(obj) && (obj.length === 0 || !isRecursiveChild))
      ? obj
      : mapObj(obj, (key, val: any) => {
        const newArray = [];

        if (Array.isArray(val)) {
          val.forEach((value) => {
            if (isObject(value) && !Array.isArray(value)) {
              newArray.push(changeStringCaseRecursive(value, type, true));
            } else {
              newArray.push(value);
            }
          });

          return [changer(key, type), newArray];
        }
        if (!val) {
          return [changer(key, type), val];
        }
        if (val instanceof Date) {
          return [changer(key, type), val];
        }
        if (isObject(val)) {
          return [changer(key, type), changeStringCaseRecursive(val, type, true)];
        }

        return [changer(key, type), val];
      });

  return Array.isArray(obj) ? objectValues(transform) : transform;
};

const booleanChanger = (objValue) => {
  if (typeof objValue == 'boolean') {
    return objValue ? 1 : 0;
  }
  return objValue;
};

export const changeBooleanToNumber = (obj: object) => {
  const transform =
    obj == null
      ? obj
      : mapObj(obj, (key, val: any) => {
        const newArray = [];
        if (Array.isArray(val)) {
          val.forEach((value) => {
            if (isObject(value) && !Array.isArray(value)) {
              newArray.push(changeBooleanToNumber(value));
            } else {
              newArray.push(booleanChanger(value));
            }
          });
          return [key, newArray];
        }
        if (!val) {
          return [key, booleanChanger(val)];
        }
        if (val instanceof Date) {
          return [key, val];
        }
        if (isObject(val)) {
          return [key, changeBooleanToNumber(val)];
        }
        return [key, booleanChanger(val)];
      });
  return Array.isArray(obj) ? objectValues(transform) : transform;
};

export default changeStringCaseRecursive;
