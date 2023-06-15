// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.local.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetSensitiveFileArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSensitiveFileArgs Empty = new GetSensitiveFileArgs();

    /**
     * Path to the file that will be read. The data source will return an error if the file does not exist.
     * 
     */
    @Import(name="filename", required=true)
    private Output<String> filename;

    /**
     * @return Path to the file that will be read. The data source will return an error if the file does not exist.
     * 
     */
    public Output<String> filename() {
        return this.filename;
    }

    private GetSensitiveFileArgs() {}

    private GetSensitiveFileArgs(GetSensitiveFileArgs $) {
        this.filename = $.filename;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSensitiveFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSensitiveFileArgs $;

        public Builder() {
            $ = new GetSensitiveFileArgs();
        }

        public Builder(GetSensitiveFileArgs defaults) {
            $ = new GetSensitiveFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param filename Path to the file that will be read. The data source will return an error if the file does not exist.
         * 
         * @return builder
         * 
         */
        public Builder filename(Output<String> filename) {
            $.filename = filename;
            return this;
        }

        /**
         * @param filename Path to the file that will be read. The data source will return an error if the file does not exist.
         * 
         * @return builder
         * 
         */
        public Builder filename(String filename) {
            return filename(Output.of(filename));
        }

        public GetSensitiveFileArgs build() {
            $.filename = Objects.requireNonNull($.filename, "expected parameter 'filename' to be non-null");
            return $;
        }
    }

}
