import defaults from 'lodash/defaults';

import {
  DataQueryRequest,
  DataQueryResponse,
  DataSourceApi,
  DataSourceInstanceSettings,
  FieldType,
  MutableDataFrame,
} from '@grafana/data';

import { defaultQuery, MyDataSourceOptions, MyQuery } from './types';

export class DataSource extends DataSourceApi<MyQuery, MyDataSourceOptions> {
  resolution: number;

  constructor(instanceSettings: DataSourceInstanceSettings<MyDataSourceOptions>) {
    super(instanceSettings);
    this.resolution = instanceSettings.jsonData.resolution || 1000.0;
  }

  async query(options: DataQueryRequest<MyQuery>): Promise<DataQueryResponse> {
    const { range } = options;
    const from = range!.from.valueOf();
    const to = range!.to.valueOf();

    // Return a constant for each query.
    const data = options.targets.map(target => {
      const query = defaults(target, defaultQuery);

      const frame = new MutableDataFrame({
        refId: query.refId,
        fields: [
          { name: 'time', type: FieldType.time },
          { name: 'value', type: FieldType.number },
        ],
      });

      const duration = to - from;
      const step = duration / this.resolution;

      for (let t = 0; t < duration; t += step) {
        frame.add({ time: from + t, value: Math.sin((2 * Math.PI * query.frequency * t) / duration) });
      }

      return frame;
    });

    return { data };
  }

  async testDatasource() {
    // Implement a health check for your data source.
    return {
      status: 'success',
      message: 'Success',
    };
  }
}
