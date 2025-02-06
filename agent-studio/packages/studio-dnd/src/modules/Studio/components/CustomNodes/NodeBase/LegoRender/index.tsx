/* eslint-disable @typescript-eslint/no-explicit-any */
import { FunctionComponent, useCallback, useMemo } from 'react';

import FormRender from '../../../DataFields/FormRender';
import Lego from '../../../Lego';
import LegoContent from '../../../LegoContent';
import TextRender from '../../../Render/TextRender';

import useStudioCategoryStore from '@/modules/Studio/stores/useStudioCategoryStore';
import useStudioDataStore from '@/modules/Studio/stores/useStudioDataStore';
import useStudioFormStore from '@/modules/Studio/stores/useStudioFormStore';
import { DataSchema, StudioCategoryOptionRenderPayload } from '@/modules/Studio/types/category';

import './LegoRender.scss';

type Props<T> = {
  background?: string;
  icon: React.ReactNode | FunctionComponent;
  id: string;
  schemaData?: DataSchema;
  title: React.ReactNode | FunctionComponent;
  idx: string;
  readonly?: boolean;
  render?: (data: StudioCategoryOptionRenderPayload<T>) => React.ReactNode;
};

const LegoRenderBase = <T,>({ background, icon, id, schemaData, title, idx, readonly }: Omit<Props<T>, 'render'>) => {
  const fields = useMemo(() => Object.keys(schemaData || {}), [schemaData]);

  const isDynamicHeight = useMemo(() => {
    if (fields.length > 1) {
      return true;
    }

    const field = fields[0];
    const fieldData = (schemaData || {})[field];

    return fieldData?.type === 'textarea';
  }, [fields, schemaData]);

  const isFixedHeight = !isDynamicHeight;

  return (
    <Lego
      background={background}
      icon={icon}
      fixedHeight={isFixedHeight}
      style={{
        width: '100%',
      }}
    >
      <LegoContent>
        <FormRender readonly={readonly} idx={idx} id={id} schemaData={schemaData}>
          <TextRender data={title} />
        </FormRender>
      </LegoContent>
    </Lego>
  );
};

const LegoRenderCustomization = <T,>({ background, id, idx, render }: Props<T>) => {
  const categoryOptionMap = useStudioCategoryStore((state) => state.categoryOptionMap);
  const allFormData = useStudioFormStore((state) => state.formMap);
  const setFormFields = useStudioFormStore((state) => state.setFormFields);
  const resetFormById = useStudioFormStore((state) => state.resetFormById);

  const data = useStudioDataStore((state) => state.data);

  const formData = allFormData[id];
  const option = categoryOptionMap[idx];

  const specifyFormFields = useCallback(
    (fields: Partial<T>) => {
      return setFormFields(id, fields as Partial<T>);
    },
    [id, setFormFields],
  );

  const resetFormData = useCallback(() => {
    return resetFormById(id);
  }, [id, resetFormById]);

  if (!render) {
    return <></>;
  }

  const childrenRender = render({
    id,
    option,
    formData: (formData || {}) as T,
    setFormFields: specifyFormFields,
    allFormData,
    data,
    resetFormData,
  });

  const legoStyle = ((childrenRender as any)?.props as any)?.['data-lego-style'] || {};
  const legoClassName = ((childrenRender as any)?.props as any)?.['data-lego-class'] || {};

  return (
    <Lego
      background={background}
      icon={undefined}
      fixedHeight={false}
      className={legoClassName}
      style={{
        width: '100%',
        ...legoStyle,
      }}
    >
      {childrenRender}
    </Lego>
  );
};

export default function LegoRender<T>({ render, ...rest }: Props<T>) {
  if (render) {
    return <LegoRenderCustomization render={render} {...rest} />;
  }

  return <LegoRenderBase {...rest} />;
}
